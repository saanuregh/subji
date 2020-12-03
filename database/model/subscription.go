package model

import (
	"time"

	"github.com/System-Glitch/goyave/v3/database"
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
