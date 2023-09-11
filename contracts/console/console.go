package console

import "github.com/urfave/cli/v2"

type Console interface {
	Call(command string)
	CallAndExit(command string)
	Engine() *cli.App
	Register(commands []Command)
	Run(arguments []string, isExitAfterCompleted bool)
}
