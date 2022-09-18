package commands

import (
	"goravel/app/models"
	"goravel/contracts/console"
	"goravel/facades"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

type Migrate struct {
}

// Signature The name and signature of the console command.
func (receiver *Migrate) Signature() string {
	return "migrate"
}

// Description The console command description.
func (receiver *Migrate) Description() string {
	return "Migrate Models"
}

// Extend The console command extend.
func (receiver *Migrate) Extend() console.CommandExtend {
	return console.CommandExtend{}
}

// Handle Execute the console command.
func (receiver *Migrate) Handle(c *cli.Context) error {
	err := facades.DB.AutoMigrate(&models.Item{})

	if err != nil {
		color.Redln("Migrate Error")
		return err
	}
	color.Greenln("Migrate success")
	return nil
}
