package subscription

import (
	"math"
	"net/http"
	"time"

	"github.com/System-Glitch/goyave/v3"
	"github.com/System-Glitch/goyave/v3/database"
	"github.com/saanuregh/subji/database/model"
	"github.com/saanuregh/subji/helper"
)

// CreateSubscriptionResponse struct for serializing CreateSubscription response.
type CreateSubscriptionResponse struct {
	Status string  `json:"status"`
	Error  string  `json:"error,omitempty"`
	Amount float64 `json:"amount"`
}

// CreateSubscription is a controller handler to create a user subscription.
func CreateSubscription(response *goyave.Response, request *goyave.Request) {
	username := request.String("user_name")
	selectedPlanID := request.String("plan_id")
	startDate := request.Date("start_date")

	availablePlans := helper.GetPlans()

	// Can't subscribe to free plan.
	if selectedPlanID == "FREE" {
		response.JSON(http.StatusInternalServerError, CreateSubscriptionResponse{
			Status: "FAILIURE",
			Error:  "Can't subscribe to free plan",
		})
		return
	}
	selectedPlan := availablePlans[selectedPlanID]

	db := database.Conn()
	user := model.User{}

	result := db.Preload("Subscriptions").First(&user, "username = ?", username)
	if response.HandleDatabaseError(result) {
		// Index to keep track of the subscription that is not free/trail and is active
		// -1 if does not exist (assuming there exist only one/zero not free/trail subscription).
		validSubscriptionIdx := -1
		for i, s := range user.Subscriptions {

			// Can't subscribe to trial plan more than once.
			if selectedPlanID == "TRIAL" && s.PlanID == "TRIAL" {
				response.JSON(http.StatusInternalServerError, CreateSubscriptionResponse{
					Status: "FAILIURE",
					Error:  "Trial exhausted",
				})
				return
			}

			if s.ValidNotFreeAndTrial() {
				validSubscriptionIdx = i
				break
			}
		}

		amount := 0.0
		if validSubscriptionIdx == -1 {
			// No valid subscription, hence add the cost of selected plan to amount.
			amount = selectedPlan.Cost
		} else {
			// Valid subscriptions exist, hence upgrade/downgrade plans based on selected plan.
			// Assuming discounts given by refunding currently active plan by remaining days left.
			validPlan := availablePlans[user.Subscriptions[validSubscriptionIdx].PlanID]
			daysLeft := user.Subscriptions[validSubscriptionIdx].DaysLeft()
			amount = selectedPlan.Cost - (daysLeft * validPlan.CostPerDay())
		}

		// Try making payment with external API.
		paymentID, err := helper.MakePayment(username, amount)
		if err != nil {
			response.JSON(http.StatusInternalServerError, CreateSubscriptionResponse{
				Status: "FAILIURE",
				Error:  "Payment failed",
			})
			return
		}

		// If payment successfull, make changes to database.
		user.Subscriptions = append(user.Subscriptions, model.Subscription{
			PlanID:    selectedPlanID,
			StartDate: startDate,
			PaymentID: paymentID,
		})
		db.Save(&user)
		// If valid subscription, make it inactive.
		if validSubscriptionIdx != -1 {
			user.Subscriptions[validSubscriptionIdx].Active = false
			db.Save(&user.Subscriptions[validSubscriptionIdx])
		}

		response.JSON(http.StatusOK, CreateSubscriptionResponse{
			Status: "SUCCESS",
			Amount: math.Round(amount*100) / 100,
		})
		return
	}

	response.JSON(http.StatusNotFound, CreateSubscriptionResponse{
		Status: "FAILIURE",
		Error:  "User not found",
	})
}

// GetSubscriptionsResponse struct for serializing GetSubscriptions response.
type GetSubscriptionsResponse struct {
	PlanID    string `json:"plan_id"`
	StartDate string `json:"start_date"`
	ValidTill string `json:"valid_till"`
}

// GetSubscriptions is a controller handler to get a user's subscriptions.
func GetSubscriptions(response *goyave.Response, request *goyave.Request) {
	user := model.User{}
	result := database.Conn().Preload("Subscriptions").First(&user, "username = ?", request.Params["username"])
	if response.HandleDatabaseError(result) {
		resp := []GetSubscriptionsResponse{}
		for _, s := range user.Subscriptions {
			validTill := s.ValidTill().Format(helper.DateLayout)
			if s.PlanID == "FREE" {
				validTill = "infinity"
			}
			if s.Active {
				resp = append(resp, GetSubscriptionsResponse{s.PlanID, s.StartDate.Format(helper.DateLayout), validTill})
			}
		}
		response.JSON(http.StatusOK, resp)
	}
}

// GetSubscriptionDateResponse struct for serializing GetSubscriptionDate response.
type GetSubscriptionDateResponse struct {
	PlanID   string      `json:"plan_id"`
	DaysLeft interface{} `json:"days_left"`
}

// GetSubscriptionDate is a controller handler to get a user's active subscription with day's left with respect to given date.
func GetSubscriptionDate(response *goyave.Response, request *goyave.Request) {
	date, err := time.Parse(helper.DateLayout, request.Params["date"])
	if err != nil {
		response.Error(err)
		return
	}
	user := model.User{}
	result := database.Conn().Preload("Subscriptions").First(&user, "username = ?", request.Params["username"])
	if response.HandleDatabaseError(result) {
		var daysLeft interface{}
		availablePlans := helper.GetPlans()
		validSubscriptionIdx := 0
		maxCost := 0.0
		daysLeft = "infinity"
		// Obtains a valid subscription with max cost and positive days left with respect to given date.
		for i, s := range user.Subscriptions {
			_daysLeft := int(helper.DaysLeft(s.ValidTill(), date))
			if availablePlans[s.PlanID].Cost >= maxCost && s.Valid() && _daysLeft > 0 {
				validSubscriptionIdx = i
				maxCost = availablePlans[s.PlanID].Cost
				if s.PlanID != "FREE" {
					daysLeft = _daysLeft
				}
			}
		}
		response.JSON(http.StatusOK, GetSubscriptionDateResponse{
			PlanID:   user.Subscriptions[validSubscriptionIdx].PlanID,
			DaysLeft: daysLeft,
		})
	}
}
