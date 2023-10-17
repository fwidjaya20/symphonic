package console

import (
	"errors"

	"github.com/fwidjaya20/symphonic/contracts/config"
	"github.com/fwidjaya20/symphonic/contracts/console"
	"github.com/golang-migrate/migrate/v4"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

type SeedCommand struct {
	config config.Config
}

func NewSeedCommand(config config.Config) console.Command {
	return &SeedCommand{
		config: config,
	}
}

func (cmd *SeedCommand) Setup() *cli.Command {
	return &cli.Command{
		Name:        "seed",
		Description: "Seed the database with records",
		Action:      cmd.Handle,
	}
}

func (cmd *SeedCommand) Handle(*cli.Context) error {
	instance, err := getSeeder(cmd.config)
	if nil != err {
		if errors.Is(err, ErrEmptyMigrationDir) {
			color.Yellowln("There is no seeder files yet.")
			return nil
		}
		return err
	}
	if nil == instance {
		color.Yellowln("Database configuration was invalid!")
		return nil
	}

	if err := instance.Up(); nil != err {
		if errors.Is(err, migrate.ErrNoChange) {
			color.Greenln("There is no new seeder files.")
			return nil
		}
		return err
	}

	color.Greenln("Seed has been completed.")

	return nil
}
