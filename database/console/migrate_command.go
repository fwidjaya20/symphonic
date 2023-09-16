package console

import (
	"errors"

	"github.com/fwidjaya20/symphonic/contracts/config"
	"github.com/fwidjaya20/symphonic/contracts/console"
	"github.com/golang-migrate/migrate/v4"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

type MigrateCommand struct {
	config config.Config
}

func NewMigrateCommand(config config.Config) console.Command {
	return &MigrateCommand{
		config: config,
	}
}

func (cmd *MigrateCommand) Setup() *cli.Command {
	return &cli.Command{
		Name:        "migrate",
		Description: "Run all database migrations",
		Action:      cmd.Handle,
	}
}

func (cmd *MigrateCommand) Handle(*cli.Context) error {
	instance, err := getMigrate(cmd.config)
	if nil != err {
		return err
	}
	if nil == instance {
		color.Yellowln("Database configuration was invalid!")
		return nil
	}

	if err := instance.Up(); nil != err {
		if errors.Is(err, migrate.ErrNoChange) {
			color.Greenln("There is no new migration files.")
			return nil
		}
		return err
	}

	color.Greenln("Migration has been completed.")

	return nil
}
