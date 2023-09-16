package database

import (
	ContractConsole "github.com/fwidjaya20/go-framework/contracts/console"
	"github.com/fwidjaya20/go-framework/contracts/foundation"
	"github.com/fwidjaya20/go-framework/database/console"
)

const Binding = "go_framework.database"

type ServiceProvider struct{}

func (provider *ServiceProvider) Boot(app foundation.Application) {}

func (provider *ServiceProvider) Register(app foundation.Application) {
	config := app.GetConfig()

	app.GetConsole().Register([]ContractConsole.Command{
		console.NewMigrateCommand(config),
		console.NewMigrateResetCommand(config),
		console.NewMigrateRollbackCommand(config),
		console.NewMigrateStatusCommand(config),
		console.NewMigrationCommand(config),
		console.NewSeedCommand(config),
		console.NewSeederCommand(config),
	})

	app.Instance(Binding, func(app foundation.Application) (any, error) {
		return &ServiceProvider{}, nil
	})
}
