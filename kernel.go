package main

import (
	"os"

	"github.com/saanuregh/subji/http/route"
	_ "github.com/saanuregh/subji/http/validation"

	"github.com/System-Glitch/goyave/v3"
	_ "github.com/System-Glitch/goyave/v3/database/dialect/postgres"
)

func main() {
	// This is the entry point of the application.
	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
}
