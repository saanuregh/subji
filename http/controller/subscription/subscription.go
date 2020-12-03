package subscription

import (
	"github.com/System-Glitch/goyave/v3"
)

// CreateSubscriptionResponse struct for serializing CreateSubscription response.
type CreateSubscriptionResponse struct {
	Status string  `json:"status"`
	Error  string  `json:"error,omitempty"`
	Amount float64 `json:"amount"`
}

// CreateSubscription is a controller handler to create a user subscription.
func CreateSubscription(response *goyave.Response, request *goyave.Request) {}

// GetSubscriptionsResponse struct for serializing GetSubscriptions response.
type GetSubscriptionsResponse struct {
	PlanID    string `json:"plan_id"`
	StartDate string `json:"start_date"`
	ValidTill string `json:"valid_till"`
}

// GetSubscriptions is a controller handler to get a user's subscriptions.
func GetSubscriptions(response *goyave.Response, request *goyave.Request) {}

// GetSubscriptionDateResponse struct for serializing GetSubscriptionDate response.
type GetSubscriptionDateResponse struct {
	PlanID   string      `json:"plan_id"`
	DaysLeft interface{} `json:"days_left"`
}

// GetSubscriptionDate is a controller handler to get a user's active subcription with day's left with respect to given date.
func GetSubscriptionDate(response *goyave.Response, request *goyave.Request) {}
