package database

import (
	"goravel/facades"
)

type ServiceProvider struct {
}

// Boot Bootstrap any application services after register.
func (database *ServiceProvider) Boot() {
}

// Register any application services.
func (database *ServiceProvider) Register() {
	app := Application{}
	facades.DB = app.Init()
}
