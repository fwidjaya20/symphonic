package console

import (
	"errors"

	ContractConfig "github.com/fwidjaya20/symphonic/contracts/config"
	ContractConsole "github.com/fwidjaya20/symphonic/contracts/console"
	"github.com/golang-migrate/migrate/v4"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

type MigrateRollbackCommand struct {
	config ContractConfig.Config
}

func NewMigrateRollbackCommand(config ContractConfig.Config) ContractConsole.Command {
	return &MigrateRollbackCommand{
		config: config,
	}
}

func (cmd *MigrateRollbackCommand) Setup() *cli.Command {
	return &cli.Command{ //nolint:exhaustruct // ignore due to cli configuration
		Name:        "migrate:rollback",
		Category:    "migrate",
		Description: "Rollback the database migrations",
		Action:      cmd.Handle,
		Flags: []cli.Flag{
			&cli.IntFlag{ //nolint:exhaustruct // ignore due to cli flag configuration
				Name:  "step",
				Usage: "rollback steps",
				Value: 1,
			},
		},
	}
}

func (cmd *MigrateRollbackCommand) Handle(ctx *cli.Context) error {
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

	if err := instance.Steps(ctx.Int("step") * -1); err != nil &&
		!errors.Is(err, migrate.ErrNoChange) &&
		!errors.Is(err, migrate.ErrNilVersion) {
		color.Redln("Migration rollback failed:", err.Error())
		return err
	}

	color.Greenln("Migrate rollback has been completed")

	return nil
}
