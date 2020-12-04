package subscription

import (
	"fmt"

	"github.com/System-Glitch/goyave/v3/validation"
	"github.com/saanuregh/subji/helper"
)

var (
	// CreateSubscriptionRequest validation rules for CreateSubscription request data.
	CreateSubscriptionRequest validation.RuleSet = validation.RuleSet{
		"user_name":  {"required", "string"},
		"plan_id":    {"required", "plan"},
		"start_date": {"required", "date", fmt.Sprint("after_equal:", helper.CurrentDateString())},
	}
)
