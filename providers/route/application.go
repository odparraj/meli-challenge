package route

import (
	"goravel/facades"
	"goravel/providers/foundation"

	"github.com/gin-gonic/gin"
)

type Application struct {
}

func (app *Application) Init() *gin.Engine {
	rootApp := foundation.Application{}

	//When running in console, don't show gin log.
	if facades.Config.GetBool("app.debug") && !rootApp.RunningInConsole() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	route := gin.New()

	return route
}
