package foundation

import (
	"goravel/contracts"
	"goravel/facades"
	"goravel/providers/config"
	"goravel/providers/console"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	configApp := config.ServiceProvider{}
	configApp.Register()

	facadesConfig := facades.Config
	facadesConfig.Add("app", map[string]interface{}{
		"providers": []contracts.ServiceProvider{
			&console.ServiceProvider{},
		},
	})

	assert.NotPanics(t, func() {
		app := Application{}
		app.Boot()
	})
}
