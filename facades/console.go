package facades

import "github.com/fwidjaya20/symphonic/contracts/console"

func Console() console.Console {
	return App().GetConsole()
}
