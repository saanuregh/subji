package helper

// Plan represents a subscribable plan.
type Plan struct {
	Validity int
	Cost     float64
}

// GetPlans gets all available plans (don't want it to be mutable).
func GetPlans() map[string]Plan {
	return map[string]Plan{
		"FREE":    {-1, 0},
		"TRIAL":   {7, 0},
		"LITE_1M": {30, 100.0},
		"PRO_1M":  {30, 200.0},
		"LITE_6M": {180, 500.0},
		"PRO_6M":  {180, 900.0},
	}
}
