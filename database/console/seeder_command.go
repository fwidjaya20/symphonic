package console

import (
	"fmt"
	"github.com/fwidjaya20/go-framework/contracts/config"
	"github.com/fwidjaya20/go-framework/contracts/console"
	"github.com/fwidjaya20/go-framework/utility/file"
	"github.com/golang-module/carbon/v2"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
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
		return err
	}

	if err := file.Create(cmd.getPath(ctx.Args().Get(0)), ""); nil != err {
		return err
	}

	color.Greenf("%s has been created.\n", cmd.getFileName(ctx.Args().Get(0)))
	color.Greenf("%s has been created.\n", cmd.getFileName(ctx.Args().Get(0)))

	return nil
}

func (cmd *SeederCommand) getFileName(name string) string {
	return fmt.Sprintf("%s_%s.sql", carbon.Now().ToShortDateTimeString(), name)
}

func (cmd *SeederCommand) getPath(name string) string {
	pwd, _ := os.Getwd()
	return fmt.Sprintf("%s/%s/seeders/%s", pwd, cmd.config.Env("database.dir", "database"), cmd.getFileName(name))
}
