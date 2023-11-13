package foundation

import (
	"github.com/fwidjaya20/symphonic/contracts/config"
	"github.com/fwidjaya20/symphonic/contracts/console"
	"github.com/fwidjaya20/symphonic/contracts/event"
	"github.com/fwidjaya20/symphonic/contracts/log"
	"github.com/fwidjaya20/symphonic/contracts/schedule"
	"gorm.io/gorm"
)

type Application interface {
	Boot()
	Get(key any) (any, error)
	GetConfig() config.Config
	GetConsole() console.Console
	GetDatabase() *gorm.DB
	GetEvent() event.Event
	GetLogger() log.Logger
	GetSchedule() schedule.Schedule
	Instance(key any, callback func(app Application) (any, error))
}
