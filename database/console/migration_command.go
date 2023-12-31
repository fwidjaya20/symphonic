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

type MigrationCommand struct {
	config ContractConfig.Config
}

func NewMigrationCommand(config ContractConfig.Config) ContractConsole.Command {
	return &MigrationCommand{
		config: config,
	}
}

func (cmd *MigrationCommand) Setup() *cli.Command {
	return &cli.Command{ //nolint:exhaustruct // ignore due to cli configuration
		Name:        "make:migration",
		Category:    "make",
		Description: "Create database migration file",
		Action:      cmd.Handle,
	}
}

func (cmd *MigrationCommand) Handle(ctx *cli.Context) error {
	if err := file.Create(cmd.getPath(ctx.Args().Get(0), "down"), ""); err != nil {
		return err
	}

	if err := file.Create(cmd.getPath(ctx.Args().Get(0), "up"), ""); err != nil {
		return err
	}

	color.Greenf("%s has been created.\n", cmd.getFileName(ctx.Args().Get(0), "down"))
	color.Greenf("%s has been created.\n", cmd.getFileName(ctx.Args().Get(0), "up"))

	return nil
}

func (cmd *MigrationCommand) getFileName(name, category string) string {
	return fmt.Sprintf("%s_%s.%s.sql", carbon.Now().ToShortDateTimeString(), name, category)
}

func (cmd *MigrationCommand) getPath(name, category string) string {
	pwd, _ := os.Getwd()

	return fmt.Sprintf(
		"%s/%s/migrations/%s",
		pwd,
		cmd.config.Get("database.dir", "database"),
		cmd.getFileName(name, category),
	)
}
