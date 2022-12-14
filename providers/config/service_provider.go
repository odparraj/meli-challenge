package config

import (
	"goravel/facades"
)

type ServiceProvider struct {
}

// Boot Bootstrap any application services after register.
func (config *ServiceProvider) Boot() {

}

// Register Register any application services.
func (config *ServiceProvider) Register() {
	app := Application{}
	facades.Config = app.Init()
}
