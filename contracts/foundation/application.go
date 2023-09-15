package foundation

import (
	"github.com/fwidjaya20/go-framework/contracts/config"
	"github.com/fwidjaya20/go-framework/contracts/console"
	"github.com/fwidjaya20/go-framework/contracts/log"
	"github.com/fwidjaya20/go-framework/contracts/schedule"
)

type Application interface {
	Boot()
	Get(key any) (any, error)
	GetConfig() config.Config
	GetConsole() console.Console
	GetLogger() log.Logger
	GetSchedule() schedule.Schedule
	Singleton(key any, callback func(app Application) (any, error))
}
