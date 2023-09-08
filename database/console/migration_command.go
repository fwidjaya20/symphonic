package console

import (
	"fmt"
	"github.com/fwidjaya20/go-framework/contracts/config"
	"github.com/fwidjaya20/go-framework/contracts/console"
	"github.com/fwidjaya20/go-framework/utility/file"
	"github.com/golang-module/carbon/v2"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

type MigrationCommand struct {
	config config.Config
}

func NewMigrationCommand(config config.Config) console.Command {
	return &MigrationCommand{
		config: config,
	}
}

func (cmd *MigrationCommand) Setup() *cli.Command {
	return &cli.Command{
		Name:        "migration",
		Description: "Create database migration file",
		Action:      cmd.Handle,
	}
}

func (cmd *MigrationCommand) Handle(ctx *cli.Context) error {
	if err := file.Create(cmd.getPath(ctx.Args().Get(0), "down"), ""); nil != err {
		log.Fatalln(err.Error())
	}

	if err := file.Create(cmd.getPath(ctx.Args().Get(0), "up"), ""); nil != err {
		log.Fatalln(err.Error())
	}

	return nil
}

func (cmd *MigrationCommand) getPath(name string, category string) string {
	pwd, _ := os.Getwd()
	return fmt.Sprintf("%s/%s/migrations/%s_%s.%s.sql", pwd, cmd.config.Env("database.dir", "database"), carbon.Now().ToShortDateTimeString(), name, category)
}
