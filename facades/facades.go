package facades

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goravel/contracts/config"
	"goravel/contracts/console"
	"goravel/contracts/http"
	"goravel/contracts/services"
	"gorm.io/gorm"
)

var Artisan console.Artisan
var Config config.Config
var DB *gorm.DB
var Log *logrus.Logger
var Request http.Request
var Response http.Response
var Route *gin.Engine
var MeliClient services.MeliApiServiceInterface
var FileImporter services.FileImportServiceInterface
var Decoders map[string]services.Decoder
