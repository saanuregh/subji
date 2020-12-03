package helper

// Plan represents a subscribable plan.
type Plan struct {
	Validity int
	Cost     float64
}

// CostPerDay calculates cost per day for the plan.
func (p *Plan) CostPerDay() float64 {
	return p.Cost / float64(p.Validity)
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
