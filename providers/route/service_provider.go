package route

import (
	"goravel/facades"
)

type ServiceProvider struct {
}

// Boot Bootstrap any application services after register.
func (route *ServiceProvider) Boot() {

}

// Register Register any application services.
func (route *ServiceProvider) Register() {
	app := Application{}
	facades.Route = app.Init()
}
