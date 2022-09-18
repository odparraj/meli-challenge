package providers

import (
	"goravel/app/console"
	"goravel/facades"
)

type ConsoleServiceProvider struct {
}

func (receiver *ConsoleServiceProvider) Boot() {

}

func (receiver *ConsoleServiceProvider) Register() {
	facades.Artisan.Register(console.Kernel{}.Commands())
}
