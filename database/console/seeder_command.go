package console

import (
	"fmt"
	"os"

	ContractConfig "github.com/fwidjaya20/symphonic/contracts/config"
	ContractConsole "github.com/fwidjaya20/symphonic/contracts/console"
	"github.com/fwidjaya20/symphonic/utility/file"
	"github.com/golang-module/carbon/v2"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

type SeederCommand struct {
	config ContractConfig.Config
}

func NewSeederCommand(config ContractConfig.Config) ContractConsole.Command {
	return &SeederCommand{
		config: config,
	}
}

func (cmd *SeederCommand) Setup() *cli.Command {
	return &cli.Command{ //nolint:exhaustruct // ignore due to cli configuration
		Name:        "make:seeder",
		Category:    "make",
		Description: "Create database seeder file",
		Action:      cmd.Handle,
	}
}

func (cmd *SeederCommand) Handle(ctx *cli.Context) error {
	if err := file.Create(cmd.getPath(ctx.Args().Get(0)), ""); err != nil {
		return err
	}

	if err := file.Create(cmd.getPath(ctx.Args().Get(0)), ""); err != nil {
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

	return fmt.Sprintf(
		"%s/%s/seeders/%s",
		pwd,
		cmd.config.Get("database.dir", "database"),
		cmd.getFileName(name),
	)
}
