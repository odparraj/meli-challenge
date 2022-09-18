package bootstrap

import (
	"goravel/config"
	"goravel/providers/foundation"
)

func Boot() {
	app := foundation.Application{}

	//Bootstrap the application
	app.Boot()

	//Bootstrap the config.
	config.Boot()
}
