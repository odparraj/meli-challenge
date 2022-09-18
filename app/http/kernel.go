package http

import (
	"github.com/gin-gonic/gin"
)

type Kernel struct {
}

func (kernel Kernel) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		gin.Logger(),
	}
}
