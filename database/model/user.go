package model

import (
	"time"

	"github.com/System-Glitch/goyave/v3/database"
	"github.com/bxcodec/faker/v3"
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

// UserGenerator generator function for User model.
func UserGenerator() interface{} {
	user := &User{}
	faker.SetGenerateUniqueValues(true)
	user.Username = faker.Name()
	user.Subscriptions = []Subscription{
		{PlanID: "FREE", StartDate: time.Now()},
	}
	return user
}
