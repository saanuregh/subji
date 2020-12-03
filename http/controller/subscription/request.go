package subscription

import (
	"fmt"
	"time"

	"github.com/System-Glitch/goyave/v3/validation"
	"github.com/saanuregh/subji/helper"
)

var (
	// CreateSubscriptionRequest validation rules for CreateSubscription request data.
	CreateSubscriptionRequest validation.RuleSet = validation.RuleSet{
		"user_name":  {"required", "string"},
		"plan_id":    {"required", "plan"},
		"start_date": {"required", "date", fmt.Sprint("after_equal:", currentDate())},
	}
)

func currentDate() string {
	t, _ := time.Parse(helper.DateLayout, time.Now().Format(helper.DateLayout))
	return t.Format("2006-01-02T15:04:05")
}
