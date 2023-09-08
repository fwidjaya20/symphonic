package console

import (
	"fmt"
	"github.com/fwidjaya20/go-framework/contracts/console"
	"github.com/urfave/cli/v2"
)

type MigrateCommand struct{}

func NewMigrateCommand() console.Command {
	return &MigrateCommand{}
}

func (cmd *MigrateCommand) Setup() *cli.Command {
	return &cli.Command{
		Name:        "migrate",
		Description: "Run all database migrations",
		Action:      cmd.Handle,
	}
}

func (cmd *MigrateCommand) Handle(*cli.Context) error {
	fmt.Println("Run all Database Migration")
	return nil
}
