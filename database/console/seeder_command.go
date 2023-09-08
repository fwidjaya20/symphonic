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

type SeederCommand struct {
	config config.Config
}

func NewSeederCommand(config config.Config) console.Command {
	return &SeederCommand{
		config: config,
	}
}

func (cmd *SeederCommand) Setup() *cli.Command {
	return &cli.Command{
		Name:        "seeder",
		Description: "Create database seeder file",
		Action:      cmd.Handle,
	}
}

func (cmd *SeederCommand) Handle(ctx *cli.Context) error {
	if err := file.Create(cmd.getPath(ctx.Args().Get(0)), ""); nil != err {
		log.Fatalln(err.Error())
	}

	if err := file.Create(cmd.getPath(ctx.Args().Get(0)), ""); nil != err {
		log.Fatalln(err.Error())
	}

	return nil
}

func (cmd *SeederCommand) getPath(name string) string {
	pwd, _ := os.Getwd()
	return fmt.Sprintf("%s/%s/seeders/%s_%s.sql", pwd, cmd.config.Env("database.dir", "database"), carbon.Now().ToShortDateTimeString(), name)
}
