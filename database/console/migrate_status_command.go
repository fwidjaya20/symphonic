package console

import (
	"github.com/fwidjaya20/symphonic/contracts/config"
	"github.com/fwidjaya20/symphonic/contracts/console"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

type MigrateStatusCommand struct {
	config config.Config
}

func NewMigrateStatusCommand(config config.Config) console.Command {
	return &MigrateStatusCommand{
		config: config,
	}
}

func (cmd *MigrateStatusCommand) Setup() *cli.Command {
	return &cli.Command{
		Name:        "migrate:status",
		Category:    "migrate",
		Description: "Show status of each migrations",
		Action:      cmd.Handle,
	}
}

func (cmd *MigrateStatusCommand) Handle(*cli.Context) error {
	instance, err := getMigrate(cmd.config)
	if nil != err {
		return err
	}
	if nil == instance {
		color.Yellowln("Database configuration was invalid!")
		return nil
	}

	version, isDirty, err := instance.Version()
	if nil != err {
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
