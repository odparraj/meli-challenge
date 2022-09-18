package console

import (
	console2 "goravel/contracts/console"
	"goravel/facades"
	"goravel/providers/console/console"
)

type ServiceProvider struct {
}

// Boot Bootstrap any application services after register.
func (receiver *ServiceProvider) Boot() {
	receiver.registerCommands()
}

// Register any application services.
func (receiver *ServiceProvider) Register() {
	app := Application{}
	facades.Artisan = app.Init()
}

func (receiver *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]console2.Command{
		&console.ConsoleMakeCommand{},
	})
}
