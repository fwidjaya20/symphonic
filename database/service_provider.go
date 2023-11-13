package database

import (
	ContractConsole "github.com/fwidjaya20/symphonic/contracts/console"
	"github.com/fwidjaya20/symphonic/contracts/foundation"
	"github.com/fwidjaya20/symphonic/database/console"
)

const Binding = "symphonic.database"

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
		return NewApplication(app.GetConfig()), nil
	})
}
