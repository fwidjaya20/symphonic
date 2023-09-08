package database

import (
	"github.com/fwidjaya20/go-framework/contracts/config"
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

func NewDatabaseServiceProvider(config config.Config) foundation.ServiceProvider {
	return &ServiceProvider{
		commands: []contractconsole.Command{
			console.NewMigrateCommand(config),
			console.NewMigrationCommand(config),
			console.NewSeederCommand(config),
		},
	}
}

func (provider *ServiceProvider) Boot() {
	app := &cli.App{Commands: factoryCommands(provider.commands)}

	if err := app.Run(os.Args); nil != err {
		log.Fatal(err)
	}
}
