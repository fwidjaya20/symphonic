package console

import "github.com/urfave/cli/v2"

type Command interface {
	Setup() *cli.Command
	Handle(ctx *cli.Context) error
}
