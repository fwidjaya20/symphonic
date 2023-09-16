package foundation

import (
	"github.com/fwidjaya20/symphonic/contracts/config"
	"github.com/fwidjaya20/symphonic/contracts/console"
	"github.com/fwidjaya20/symphonic/contracts/log"
	"github.com/fwidjaya20/symphonic/contracts/schedule"
)

type Application interface {
	Boot()
	Get(key any) (any, error)
	GetConfig() config.Config
	GetConsole() console.Console
	GetLogger() log.Logger
	GetSchedule() schedule.Schedule
	Instance(key any, callback func(app Application) (any, error))
}
