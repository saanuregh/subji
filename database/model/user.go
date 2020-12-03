package model

import (
	"github.com/System-Glitch/goyave/v3/database"
	"gorm.io/gorm"
)

func init() {
	database.RegisterModel(&User{})
}

// User represents an user.
type User struct {
	gorm.Model
	// Username are assumed to be unique
	Username      string `gorm:"type:char(100);uniqueIndex"`
	Subscriptions []Subscription
}
