package console

import (
	"errors"
	"github.com/fwidjaya20/go-framework/contracts/config"
	"github.com/fwidjaya20/go-framework/contracts/console"
	"github.com/golang-migrate/migrate/v4"
	"github.com/urfave/cli/v2"
	"log"
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
		log.Fatalln("Database config was invalid!")
		return nil
	}

	if err := instance.Up(); nil != err && errors.Is(err, migrate.ErrNoChange) {
		log.Fatalln("Database Migrate has been failed!", err.Error())
		return err
	}

	return nil
}
