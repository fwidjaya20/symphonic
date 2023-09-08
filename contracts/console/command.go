package console

import "github.com/urfave/cli/v2"

type Command interface {
	Signature() string
	Description() string
	Extend() cli.Command
	Handle(ctx *cli.Context) error
}
