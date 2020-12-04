package route

import (
	"net/http"

	"github.com/saanuregh/subji/http/controller/subscription"
	"github.com/saanuregh/subji/http/controller/user"

	"github.com/System-Glitch/goyave/v3"
	"github.com/System-Glitch/goyave/v3/config"
	"github.com/System-Glitch/goyave/v3/cors"
)

// Register all the application routes. This is the main route registrer.
func Register(router *goyave.Router) {

	// Applying default CORS settings (allow all methods and all origins).
	router.CORS(cors.Default())

	// User routes.
	router.Get("/user/{username:[a-zA-Z0-9]+}", user.GetUser)
	router.Put("/user/{username:[a-zA-Z0-9]+}", user.CreateUser)

	// Subscription routes.
	router.Post("/subscription", subscription.CreateSubscription).Validate(subscription.CreateSubscriptionRequest)
	router.Get("/subscription/{username:[a-zA-Z0-9]+}", subscription.GetSubscriptions)
	router.Get("/subscription/{username:[a-zA-Z0-9]+}/{date:[0-9]{4}-[0-9]{2}-[0-9]{2}}", subscription.GetSubscriptionDate)

	// Mock payment route for testing purposes.
	if config.GetString("app.environment") == "test" {
		router.Post("/payment", func(response *goyave.Response, request *goyave.Request) {
			response.JSON(http.StatusOK, map[string]string{
				"status":     "SUCCESS",
				"payment_id": "24242-3443-sdstg-3343",
			})
		})
	}
}
