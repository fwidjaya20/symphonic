package database

import (
	"github.com/fwidjaya20/go-framework/contracts/console"
	"github.com/urfave/cli/v2"
)

func factoryCommands(commands []console.Command) []*cli.Command {
	var objCmd = make([]*cli.Command, len(commands))

	for i, it := range commands {
		objCmd[i] = it.Setup()
	}

	return objCmd
}
