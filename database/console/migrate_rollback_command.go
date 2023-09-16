package console

import (
	"errors"

	"github.com/fwidjaya20/symphonic/contracts/config"
	"github.com/fwidjaya20/symphonic/contracts/console"
	"github.com/golang-migrate/migrate/v4"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

type MigrateRollbackCommand struct {
	config config.Config
}

func NewMigrateRollbackCommand(config config.Config) console.Command {
	return &MigrateRollbackCommand{
		config: config,
	}
}

func (cmd *MigrateRollbackCommand) Setup() *cli.Command {
	return &cli.Command{
		Name:        "migrate:rollback",
		Category:    "migrate",
		Description: "Rollback the database migrations",
		Action:      cmd.Handle,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "step",
				Usage: "rollback steps",
				Value: 1,
			},
		},
	}
}

func (cmd *MigrateRollbackCommand) Handle(ctx *cli.Context) error {
	instance, err := getMigrate(cmd.config)
	if nil != err {
		return err
	}
	if nil == instance {
		color.Yellowln("Database configuration was invalid!")
		return nil
	}

	if err := instance.Steps(ctx.Int("step") * -1); nil != err && !errors.Is(err, migrate.ErrNoChange) && !errors.Is(err, migrate.ErrNilVersion) {
		color.Redln("Migration rollback failed:", err.Error())
		return err
	}

	color.Greenln("Migrate rollback has been completed")

	return nil
}
