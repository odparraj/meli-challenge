package console

import (
	"os"
	"testing"

	"goravel/contracts"
	console2 "goravel/contracts/console"
	"goravel/facades"
	"goravel/providers/config"
	"goravel/providers/console"
	testing2 "goravel/providers/support/testing"

	"github.com/stretchr/testify/assert"
)

func TestListCommand(t *testing.T) {
	err := testing2.CreateEnv()
	assert.Nil(t, err)

	configApp := config.ServiceProvider{}
	configApp.Register()

	facadesConfig := facades.Config
	facadesConfig.Add("app", map[string]interface{}{
		"providers": []contracts.ServiceProvider{},
	})

	consoleApp := console.Application{}
	facades.Artisan = consoleApp.Init()
	consoleApp.Init().Register([]console2.Command{
		&ListCommand{},
	})

	assert.NotPanics(t, func() {
		consoleApp.Call("list")
	})

	err = os.Remove(".env")
	assert.Nil(t, err)
}
