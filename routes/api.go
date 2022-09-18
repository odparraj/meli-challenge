package routes

import (
	"goravel/app/http/controllers"
	"goravel/facades"
)

func Api() {
	facades.Route.POST("proccess-file", controllers.ProccessFileController{}.Handle)
}
