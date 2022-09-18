package providers

import (
	"goravel/app/http"
	"goravel/facades"
	"goravel/routes"
)

type RouteServiceProvider struct {
}

func (receiver *RouteServiceProvider) Boot() {

}

func (receiver *RouteServiceProvider) Register() {
	//Add HTTP middlewares.
	facades.Route.Use(http.Kernel{}.Middleware()...)

	//Add routes
	routes.Api()
}
