package database

import (
	"testing"

	"goravel/providers/config"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	configApp := config.ServiceProvider{}
	configApp.Register()

	assert.NotPanics(t, func() {
		app := Application{}
		app.Init()
	})
}
