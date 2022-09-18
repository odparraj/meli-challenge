package http

import (
	"goravel/contracts/http"

	"github.com/gin-gonic/gin"
)

type Application struct {
}

func (app *Application) Init() (http.Request, http.Response) {
	return request{}, response{}
}

type request struct {
}

func (r request) Validate(ctx *gin.Context, request http.FormRequest) []error {
	if err := ctx.ShouldBind(request); err != nil {
	}

	return nil
}

type response struct {
}

func (r response) Success(ctx *gin.Context, data interface{}) {
	r.Custom(ctx, data, 200)
}

func (r response) Custom(ctx *gin.Context, data interface{}, code int) {
	ctx.JSON(code, data)
}
