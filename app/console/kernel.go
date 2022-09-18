package console

import (
	"goravel/app/console/commands"
	"goravel/contracts/console"
)

type Kernel struct {
}

func (kernel Kernel) Commands() []console.Command {
	return []console.Command{&commands.Migrate{}}
}
