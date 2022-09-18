package providers

import (
	console2 "goravel/contracts/console"
	"goravel/facades"
	"goravel/providers/foundation/console"
)

type ArtisanServiceProvider struct {
}

// Boot Bootstrap any application services after register.
func (artisan *ArtisanServiceProvider) Boot() {
	artisan.registerCommands()
}

// Register any application services.
func (artisan *ArtisanServiceProvider) Register() {

}

func (artisan *ArtisanServiceProvider) registerCommands() {
	facades.Artisan.Register([]console2.Command{
		&console.ListCommand{},
	})
}
