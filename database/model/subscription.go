package model

import (
	"time"

	"github.com/System-Glitch/goyave/v3/database"
	"github.com/saanuregh/subji/helper"
	"gorm.io/gorm"
)

func init() {
	database.RegisterModel(&Subscription{})
}

// Subscription represents an user subscription.
type Subscription struct {
	gorm.Model
	UserID    uint
	PlanID    string
	StartDate time.Time
	// Active indicates if the subscription is cancelled or not.
	Active    bool `gorm:"default:true"`
	PaymentID string
}

// ValidTill returns date till the subscription is valid.
func (s *Subscription) ValidTill() time.Time {
	plan := helper.GetPlans()[s.PlanID]
	return s.StartDate.Add(time.Duration(plan.Validity*24) * time.Hour)
}

// Valid checks if the subscription is still valid.
func (s *Subscription) Valid() bool {
	now := time.Now()
	if s.Active && s.ValidTill().After(now) {
		return true
	}
	return false
}

// ValidNotFree checks if the subscription is still valid excluding free plan.
func (s *Subscription) ValidNotFree() bool {
	if s.PlanID != "FREE" && s.Valid() {
		return true
	}
	return false
}

// ValidNotFreeAndTrial checks if the subscription is still valid excluding free and trial plan.
func (s *Subscription) ValidNotFreeAndTrial() bool {
	if s.PlanID != "FREE" && s.PlanID != "TRIAL" && s.Valid() {
		return true
	}
	return false
}

// DaysLeft returns number of days left in subscription.
func (s *Subscription) DaysLeft() float64 {
	now := time.Now()
	return helper.DaysLeft(s.ValidTill(), now)
}
