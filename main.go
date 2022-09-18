package main

import (
	"goravel/bootstrap"
	"goravel/facades"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	if err := facades.Route.Run(facades.Config.GetString("app.host")); err != nil {
		facades.Log.Errorf("Route run error: %v", err)
	}
}
