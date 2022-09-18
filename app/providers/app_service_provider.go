package providers

import (
	"goravel/app/services"
	s "goravel/contracts/services"
	"goravel/facades"
)

type AppServiceProvider struct {
}

func (receiver *AppServiceProvider) Boot() {
}

func (receiver *AppServiceProvider) Register() {
	facades.MeliClient = &services.MeliApiService{}
	facades.FileImporter = &services.FileImportService{}

	var decoders = make(map[string]s.Decoder)

	decoders["jsonln"] = &services.JsonlnDecoder{}
	decoders["csv"] = &services.CsvDecoder{}
	decoders["txt"] = &services.CsvDecoder{}

	facades.Decoders = decoders
}
