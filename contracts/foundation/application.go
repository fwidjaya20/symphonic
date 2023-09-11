package foundation

import (
	"github.com/fwidjaya20/go-framework/contracts/config"
	"github.com/fwidjaya20/go-framework/contracts/console"
)

type Application interface {
	Boot()
	Get(key any) (any, error)
	GetConfig() config.Config
	GetConsole() console.Console
	Singleton(key any, callback func(app Application) (any, error))
}
