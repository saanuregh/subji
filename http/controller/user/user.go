package user

import (
	"net/http"
	"strings"
	"time"

	"github.com/System-Glitch/goyave/v3"
	"github.com/System-Glitch/goyave/v3/database"
	"github.com/saanuregh/subji/database/model"
	"github.com/saanuregh/subji/helper"
)

// CreateUser is a controller handler to create a new user.
func CreateUser(response *goyave.Response, request *goyave.Request) {
	user := model.User{
		Username: request.Params["username"],
		Subscriptions: []model.Subscription{
			{PlanID: "FREE", StartDate: time.Now()},
		},
	}
	result := database.Conn().Create(&user)
	if response.HandleDatabaseError(result) {
		response.Status(http.StatusOK)
	}
}

// GetUserResponse struct for serializing GetUser response.
type GetUserResponse struct {
	Username  string `json:"user_name"`
	CreatedAt string `json:"created_at"`
}

// GetUser is a controller handler to request an user with created date.
func GetUser(response *goyave.Response, request *goyave.Request) {
	user := model.User{}
	result := database.Conn().First(&user, "username = ?", request.Params["username"])
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, GetUserResponse{
			Username:  strings.TrimSpace(user.Username),
			CreatedAt: user.CreatedAt.Format(helper.DateTimeLayout),
		})
	}
}
