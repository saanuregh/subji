package test

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/System-Glitch/goyave/v3"
	"github.com/System-Glitch/goyave/v3/database"
	_ "github.com/System-Glitch/goyave/v3/database/dialect/postgres"
	"github.com/saanuregh/subji/database/model"
	"github.com/saanuregh/subji/helper"
	"github.com/saanuregh/subji/http/controller/subscription"
	"github.com/saanuregh/subji/http/route"
	_ "github.com/saanuregh/subji/http/validation"
)

// Test suite for the Subscription controller.
type SubscriptionTestSuite struct {
	goyave.TestSuite
}

func (suite *SubscriptionTestSuite) SetupTest() {
	suite.ClearDatabase()
}

// TestCreateSubscription tests POST "/subscription" route.
func (suite *SubscriptionTestSuite) TestCreateSubscription() {
	suite.RunServer(route.Register, func() {
		username := "subtest1"
		_ = database.NewFactory(model.UserGenerator).Override(&model.User{
			Username: username,
		}).Save(1).([]*model.User)
		headers := map[string]string{"Content-Type": "application/json"}
		body, _ := json.Marshal(map[string]interface{}{
			"user_name":  username,
			"plan_id":    "LITE_1M",
			"start_date": time.Now().Format(helper.DateLayout),
		})
		resp, err := suite.Post("/subscription", headers, bytes.NewReader(body))
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
			json := subscription.CreateSubscriptionResponse{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			suite.NotNil(json)
			if err == nil {
				suite.Equal("SUCCESS", json.Status)
				suite.Equal(100.0, json.Amount)
			}
		}
	})
}

// TestCreateSubscriptionUpgrade tests POST "/subscription" route (upgrade case).
func (suite *SubscriptionTestSuite) TestCreateSubscriptionUpgrade() {
	suite.RunServer(route.Register, func() {
		username := "subtest2"
		_ = database.NewFactory(model.UserGenerator).Override(&model.User{
			Username: username,
			Subscriptions: []model.Subscription{
				{PlanID: "FREE", StartDate: time.Now()},
				{PlanID: "LITE_1M", StartDate: time.Now()},
			},
		}).Save(1).([]*model.User)
		headers := map[string]string{"Content-Type": "application/json"}
		body, _ := json.Marshal(map[string]interface{}{
			"user_name":  username,
			"plan_id":    "PRO_6M",
			"start_date": time.Now().Format(helper.DateLayout),
		})
		resp, err := suite.Post("/subscription", headers, bytes.NewReader(body))
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
			json := subscription.CreateSubscriptionResponse{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			suite.NotNil(json)
			suite.NotNil(json)
			if err == nil {
				suite.Equal("SUCCESS", json.Status)
				suite.Equal(800, int(json.Amount))
			}
		}
	})
}

// TestCreateSubscriptionDowngrade tests POST "/subscription" route (downgrade case).
func (suite *SubscriptionTestSuite) TestCreateSubscriptionDowngrade() {
	suite.RunServer(route.Register, func() {
		username := "subtest3"
		_ = database.NewFactory(model.UserGenerator).Override(&model.User{
			Username: username,
			Subscriptions: []model.Subscription{
				{PlanID: "FREE", StartDate: time.Now()},
				{PlanID: "PRO_6M", StartDate: time.Now()},
			},
		}).Save(1).([]*model.User)
		headers := map[string]string{"Content-Type": "application/json"}
		body, _ := json.Marshal(map[string]interface{}{
			"user_name":  username,
			"plan_id":    "LITE_1M",
			"start_date": time.Now().Format(helper.DateLayout),
		})
		resp, err := suite.Post("/subscription", headers, bytes.NewReader(body))
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
			json := subscription.CreateSubscriptionResponse{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			suite.NotNil(json)
			if err == nil {
				suite.Equal("SUCCESS", json.Status)
				suite.Equal(-800, int(json.Amount))
			}
		}
	})
}

// TestCreateSubscriptionUserDontExist tests POST "/subscription" route (user doesn't exit case).
func (suite *SubscriptionTestSuite) TestCreateSubscriptionUserDontExist() {
	suite.RunServer(route.Register, func() {
		username := "subtestdontexist"
		headers := map[string]string{"Content-Type": "application/json"}
		body, _ := json.Marshal(map[string]interface{}{
			"user_name":  username,
			"plan_id":    "LITE_1M",
			"start_date": time.Now().Format(helper.DateLayout),
		})
		resp, err := suite.Post("/subscription", headers, bytes.NewReader(body))
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(404, resp.StatusCode)
			json := subscription.CreateSubscriptionResponse{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			suite.NotNil(json)
			suite.NotNil(json)
			if err == nil {
				suite.Equal("FAILIURE", json.Status)
				suite.Equal(0.0, json.Amount)
			}
		}
	})
}

// TestCreateSubscriptionFreePlan tests POST "/subscription" route (free plan case).
func (suite *SubscriptionTestSuite) TestCreateSubscriptionFreePlan() {
	suite.RunServer(route.Register, func() {
		username := "subtest4"
		_ = database.NewFactory(model.UserGenerator).Override(&model.User{
			Username: username,
		}).Save(1).([]*model.User)
		headers := map[string]string{"Content-Type": "application/json"}
		body, _ := json.Marshal(map[string]interface{}{
			"user_name":  username,
			"plan_id":    "FREE",
			"start_date": time.Now().Format(helper.DateLayout),
		})
		resp, err := suite.Post("/subscription", headers, bytes.NewReader(body))
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(500, resp.StatusCode)
			json := subscription.CreateSubscriptionResponse{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			suite.NotNil(json)
			if err == nil {
				suite.Equal("FAILIURE", json.Status)
				suite.Equal(0.0, json.Amount)
			}
		}
	})
}

// TestGetSubscription tests GET "/subscription/{username}" route.
func (suite *SubscriptionTestSuite) TestGetSubscription() {
	suite.RunServer(route.Register, func() {
		username := "subtest5"
		subscriptions := []model.Subscription{
			{PlanID: "FREE", StartDate: time.Now()},
			{PlanID: "PRO_6M", StartDate: time.Now()},
		}
		_ = database.NewFactory(model.UserGenerator).Override(&model.User{
			Username:      username,
			Subscriptions: subscriptions,
		}).Save(1).([]*model.User)
		resp, err := suite.Get("/subscription/"+username, nil)
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
			json := []subscription.GetSubscriptionsResponse{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			suite.NotNil(json)
			if err == nil {
				now := time.Now()
				suite.Equal(len(json), 2)
				suite.Equal(json[0].PlanID, "FREE")
				suite.Equal(json[0].StartDate, now.Format(helper.DateLayout))
				suite.Equal(json[0].ValidTill, "infinity")
				suite.Equal(json[1].PlanID, "PRO_6M")
				suite.Equal(json[1].StartDate, now.Format(helper.DateLayout))
				suite.Equal(json[1].ValidTill, subscriptions[1].ValidTill().Format(helper.DateLayout))
			}
		}
	})
}

// TestGetSubscriptionDate tests GET "/subscription/{username}/{date}" route.
func (suite *SubscriptionTestSuite) TestGetSubscriptionDate() {
	suite.RunServer(route.Register, func() {
		username := "subtest6"
		dateString := "2021-01-04"
		date, _ := time.Parse(helper.DateLayout, dateString)
		subscriptions := []model.Subscription{
			{PlanID: "FREE", StartDate: time.Now()},
			{PlanID: "PRO_6M", StartDate: time.Now()},
		}
		_ = database.NewFactory(model.UserGenerator).Override(&model.User{
			Username:      username,
			Subscriptions: subscriptions,
		}).Save(1).([]*model.User)
		resp, err := suite.Get("/subscription/"+username+"/"+dateString, nil)
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
			json := subscription.GetSubscriptionDateResponse{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			suite.NotNil(json)
			if err == nil {
				suite.Equal(json.PlanID, "PRO_6M")
				suite.Equal(int(json.DaysLeft.(float64)), int(subscriptions[1].ValidTill().Sub(date).Hours()/24))
			}
		}
	})
}

// TestMakePayment tests POST "/payment" route (only for testing purpose).
func (suite *SubscriptionTestSuite) TestMakePayment() {
	suite.RunServer(route.Register, func() {
		_, err := helper.MakePayment("test", 100)
		suite.Nil(err)
	})
}

func TestSubscriptionSuite(t *testing.T) {
	goyave.RunTest(t, new(SubscriptionTestSuite))
}
