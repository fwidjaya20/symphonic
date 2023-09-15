package facades

import "github.com/fwidjaya20/go-framework/contracts/log"

func Logger() log.Logger {
	return App().GetLogger()
}
