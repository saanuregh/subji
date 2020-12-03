package test

import (
	"testing"
	"time"

	"github.com/System-Glitch/goyave/v3"
	"github.com/System-Glitch/goyave/v3/database"
	_ "github.com/System-Glitch/goyave/v3/database/dialect/postgres"
	"github.com/saanuregh/subji/database/model"
	"github.com/saanuregh/subji/helper"
	"github.com/saanuregh/subji/http/controller/user"
	"github.com/saanuregh/subji/http/route"
	_ "github.com/saanuregh/subji/http/validation"
)

// Test suite for the User controller.
type UserTestSuite struct {
	goyave.TestSuite
}

func (suite *UserTestSuite) SetupTest() {
	suite.ClearDatabase()
}

// TestCreateUser tests PUT "/user/{username}" route.
func (suite *UserTestSuite) TestCreateUser() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Put("/user/test1", nil, nil)
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
		}
	})
}

// TestCreateUserDuplicate tests PUT "/user/{username}" route (duplicate username case).
func (suite *UserTestSuite) TestCreateUserDuplicate() {
	suite.RunServer(route.Register, func() {
		username := "test2"
		_ = database.NewFactory(model.UserGenerator).Override(&model.User{
			Username: username,
		}).Save(1).([]*model.User)
		resp, err := suite.Put("/user/"+username, nil, nil)
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(500, resp.StatusCode)
		}
	})
}

// TestGetUser tests GET "/user/{username}" route.
func (suite *UserTestSuite) TestGetUser() {
	suite.RunServer(route.Register, func() {
		username := "test3"
		_ = database.NewFactory(model.UserGenerator).Override(&model.User{
			Username: username,
		}).Save(1).([]*model.User)
		resp, err := suite.Get("/user/"+username, nil)
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
			json := user.GetUserResponse{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			suite.NotNil(json)
			if err == nil {
				suite.Equal(username, json.Username)
				_, err = time.Parse(helper.DateTimeLayout, json.CreatedAt)
				suite.Nil(err)
			}
		}
	})
}

// TestGetUserDontExist tests PUT "/user/{username}" route (user doesn't exist case).
func (suite *UserTestSuite) TestGetUserDontExist() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/user/dontexist", nil)
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(404, resp.StatusCode)
		}
	})
}

func TestUserSuite(t *testing.T) {
	goyave.RunTest(t, new(UserTestSuite))
}
