package user

import (
	"github.com/System-Glitch/goyave/v3"
)

// CreateUser is a controller handler to create a new user.
func CreateUser(response *goyave.Response, request *goyave.Request) {}

// GetUserResponse struct for serializing GetUser response.
type GetUserResponse struct {
	Username  string `json:"user_name"`
	CreatedAt string `json:"created_at"`
}

// GetUser is a controller handler to request an user with created date.
func GetUser(response *goyave.Response, request *goyave.Request) {}
