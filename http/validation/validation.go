package validation

import (
	"github.com/System-Glitch/goyave/v3/validation"
	"github.com/saanuregh/subji/helper"
)

// validateSubscriptionPlan custom validation rule for PlanID.
func validateSubscriptionPlan(field string, value interface{}, parameters []string, form map[string]interface{}) bool {
	if str, ok := value.(string); ok {
		if _, exist := helper.GetPlans()[str]; exist {
			return true
		}
	}
	return false
}

func init() {
	validation.AddRule("plan", &validation.RuleDefinition{
		Function:           validateSubscriptionPlan,
		RequiredParameters: 0,
	})
}
