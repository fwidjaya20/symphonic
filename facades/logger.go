package facades

import "github.com/fwidjaya20/symphonic/contracts/log"

func Logger() log.Logger {
	return App().GetLogger()
}
