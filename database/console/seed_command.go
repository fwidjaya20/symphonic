//nolint:dupl // ignore due to database migrate logic
package console

import (
	"errors"

	ContractConfig "github.com/fwidjaya20/symphonic/contracts/config"
	ContractConsole "github.com/fwidjaya20/symphonic/contracts/console"
	"github.com/golang-migrate/migrate/v4"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

type SeedCommand struct {
	config ContractConfig.Config
}

func NewSeedCommand(config ContractConfig.Config) ContractConsole.Command {
	return &SeedCommand{
		config: config,
	}
}

func (cmd *SeedCommand) Setup() *cli.Command {
	return &cli.Command{ //nolint:exhaustruct // ignore due to cli configuration
		Name:        "seed",
		Description: "Seed the database with records",
		Action:      cmd.Handle,
	}
}

func (cmd *SeedCommand) Handle(*cli.Context) error {
	instance, err := getSeeder(cmd.config)
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

	if err := instance.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			color.Greenln("There is no new seeder files.")
			return nil
		}

		return err
	}

	color.Greenln("Seed has been completed.")

	return nil
}
