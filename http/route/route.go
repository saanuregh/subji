package route

import (
	"github.com/saanuregh/subji/http/controller/subscription"
	"github.com/saanuregh/subji/http/controller/user"

	"github.com/System-Glitch/goyave/v3"
	"github.com/System-Glitch/goyave/v3/cors"
)

// Register all the application routes. This is the main route registrer.
func Register(router *goyave.Router) {

	// Applying default CORS settings (allow all methods and all origins).
	router.CORS(cors.Default())

	// User routes.
	router.Get("/user/{username}", user.GetUser)
	router.Put("/user/{username}", user.CreateUser)

	// Subscription routes.
	router.Post("/subscription", subscription.CreateSubscription)
	router.Get("/subscription/{username}", subscription.GetSubscriptions)
	router.Get("/subscription/{username}/{date}", subscription.GetSubscriptionDate)
}
