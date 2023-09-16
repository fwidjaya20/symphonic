package console

import (
	"errors"

	"github.com/fwidjaya20/symphonic/contracts/config"
	"github.com/fwidjaya20/symphonic/contracts/console"
	"github.com/golang-migrate/migrate/v4"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

type MigrateResetCommand struct {
	config config.Config
}

func NewMigrateResetCommand(config config.Config) console.Command {
	return &MigrateResetCommand{
		config: config,
	}
}

func (cmd *MigrateResetCommand) Setup() *cli.Command {
	return &cli.Command{
		Name:        "migrate:reset",
		Category:    "migrate",
		Description: "Rollback all database migrations",
		Action:      cmd.Handle,
	}
}

func (cmd *MigrateResetCommand) Handle(*cli.Context) error {
	instance, err := getMigrate(cmd.config)
	if nil != err {
		return err
	}
	if nil == instance {
		color.Yellowln("Database configuration was invalid!")
		return nil
	}

	if err := instance.Down(); nil != err && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	instance, err = getSeeder(cmd.config)
	if nil != err {
		return err
	}
	if nil == instance {
		color.Yellowln("Database configuration was invalid!")
		return nil
	}

	if err := instance.Down(); nil != err && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	color.Greenln("Migration reset has been completed.")

	return nil
}
