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

func (cmd *MigrateCommand) Signature() string {
	return "migrate"
}

func (cmd *MigrateCommand) Description() string {
	return "Run all database migrations"
}

func (cmd *MigrateCommand) Extend() cli.Command {
	return cli.Command{}
}

func (cmd *MigrateCommand) Handle(*cli.Context) error {
	fmt.Println("Hello World")
	return nil
}
