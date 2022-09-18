package controllers

import (
	"goravel/facades"
	"goravel/providers/support"
	"github.com/gin-gonic/gin"
)

type ProccessFileRequest struct {
	File      string `json:"file" binding:"required"`
	Decoder   string `json:"decoder" binding:"required,oneof=txt jsonln csv"`
	Delimiter string `json:"delimiter"`
}

type ProccessFileController struct {
}

func (r ProccessFileController) Handle(ctx *gin.Context) {

	requestBody := ProccessFileRequest{}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		facades.Response.Custom(
			ctx,
			gin.H{"message": err.Error()},
			422,
		)
		return
	}

	facades.Config.Add("decoder", map[string]interface{}{
		"type":      requestBody.Decoder,
		"delimiter": requestBody.Delimiter,
	})

	status, errors := facades.FileImporter.ProcessFile(requestBody.File)

	if errors != nil {
		facades.Response.Success(
			ctx,
			gin.H{
				"message": status,
				"errors": support.GetErrorMessages(errors),
			},
		)
		return
	}

	facades.Response.Success(
		ctx,
		gin.H{"message": status},
	)
}
