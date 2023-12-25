package console

import (
	"os"

	"github.com/fwidjaya20/symphonic/contracts/console"
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
	artisanIndex := -1

	for i, it := range arguments {
		if it == "artisan" {
			artisanIndex = i
			break
		}
	}

	if artisanIndex == -1 {
		return
	}

	if err := a.Engine().Run(append([]string{arguments[0]}, arguments[artisanIndex+1:]...)); err != nil {
		panic(err.Error())
	}

	if isExitAfterComplete {
		os.Exit(0)
	}
}
