package database

import (
	contractconsole "github.com/fwidjaya20/go-framework/contracts/console"
	"github.com/fwidjaya20/go-framework/contracts/foundation"
	"github.com/fwidjaya20/go-framework/database/console"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

type ServiceProvider struct {
	commands []contractconsole.Command
}

func NewDatabaseServiceProvider() foundation.ServiceProvider {
	return &ServiceProvider{
		commands: []contractconsole.Command{
			console.NewMigrateCommand(),
		},
	}
}

func (provider *ServiceProvider) Boot() {
	app := &cli.App{Commands: factoryCommands(provider.commands)}

	if err := app.Run(os.Args); nil != err {
		log.Fatal(err)
	}
}
