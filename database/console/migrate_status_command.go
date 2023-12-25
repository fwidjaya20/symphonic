package console

import (
	"errors"

	ContractConfig "github.com/fwidjaya20/symphonic/contracts/config"
	ContractConsole "github.com/fwidjaya20/symphonic/contracts/console"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

type MigrateStatusCommand struct {
	config ContractConfig.Config
}

func NewMigrateStatusCommand(config ContractConfig.Config) ContractConsole.Command {
	return &MigrateStatusCommand{
		config: config,
	}
}

func (cmd *MigrateStatusCommand) Setup() *cli.Command {
	return &cli.Command{ //nolint:exhaustruct // ignore due to cli configuration
		Name:        "migrate:status",
		Category:    "migrate",
		Description: "Show status of each migrations",
		Action:      cmd.Handle,
	}
}

func (cmd *MigrateStatusCommand) Handle(*cli.Context) error {
	instance, err := getMigrate(cmd.config)
	if err != nil {
		if errors.Is(err, ErrEmptyMigrationDir) {
			color.Yellowln("There is no seeder files yet.")
			return nil
		}

		return err
	}

	if instance == nil {
		color.Yellowln("Database configuration was invalid!")
		return nil
	}

	version, isDirty, err := instance.Version()
	if err != nil {
		return err
	}

	if isDirty {
		color.Yellowln("Migration Status: DIRTY")
	} else {
		color.Yellowln("Migration Status: CLEAN")
	}

	color.Greenf("Migration Version: %d\n", version)

	return nil
}
