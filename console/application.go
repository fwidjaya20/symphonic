package console

import (
	"os"

	"github.com/fwidjaya20/go-framework/contracts/console"
	"github.com/urfave/cli/v2"
)

type Application struct {
	cli *cli.App
}

func NewApplication() console.Console {
	instance := cli.NewApp()

	return &Application{cli: instance}
}

func (a *Application) Call(command string) {
	a.Run(append(os.Args, command), false)
}

func (a *Application) CallAndExit(command string) {
	a.Run(append(os.Args, command), true)
}

func (a *Application) Engine() *cli.App {
	return a.cli
}

func (a *Application) Register(commands []console.Command) {
	for _, it := range commands {
		a.cli.Commands = append(a.cli.Commands, it.Setup())
	}
}

func (a *Application) Run(arguments []string, isExitAfterComplete bool) {
	if err := a.Engine().Run(arguments); nil != err {
		panic(err.Error())
	}

	if isExitAfterComplete {
		os.Exit(0)
	}
}
