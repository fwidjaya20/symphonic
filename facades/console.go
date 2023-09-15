package facades

import "github.com/fwidjaya20/go-framework/contracts/console"

func Console() console.Console {
	return App().GetConsole()
}
