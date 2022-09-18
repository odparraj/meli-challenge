package config

import (
	"goravel/app/providers"
	"goravel/contracts"
	"goravel/facades"
	"goravel/providers/console"
	"goravel/providers/database"
	foundation "goravel/providers/foundation/providers"
	"goravel/providers/http"
	"goravel/providers/route"
)

// Boot Start all init methods of the current folder to bootstrap all config.
func Boot() {}

func init() {
	config := facades.Config
	config.Add("app", map[string]interface{}{
		//Application Name
		//This value is the name of your application. This value is used when the
		//framework needs to place the application's name in a notification or
		//any other location as required by the application or its packages.
		"name": config.Env("APP_NAME", "nft"),

		//Application Environment
		//This value determines the "environment" your application is currently
		//running in. This may determine how you prefer to configure various
		//services the application utilizes. Set this in your ".env" file.
		"env": config.Env("APP_ENV", "production"),

		//Application Debug Mode
		"debug": config.Env("APP_DEBUG", false),

		//Application host, http server listening address.
		"host": config.Env("APP_HOST", "127.0.0.1:3000"),

		//Autoload service providers
		//The service providers listed here will be automatically loaded on the
		//request to your application. Feel free to add your own services to
		//this array to grant expanded functionality to your applications.
		"providers": []contracts.ServiceProvider{
			&console.ServiceProvider{},
			&database.ServiceProvider{},
			&http.ServiceProvider{},
			&foundation.ArtisanServiceProvider{},
			&route.ServiceProvider{},
			&providers.AppServiceProvider{},
			&providers.RouteServiceProvider{},
			&providers.ConsoleServiceProvider{},
		},
	})
}
