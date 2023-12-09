package schedule

import (
	"github.com/fwidjaya20/symphonic/contracts/log"
	"github.com/fwidjaya20/symphonic/contracts/schedule"
	"github.com/robfig/cron/v3"
)

type Application struct {
	cron   *cron.Cron
	logger log.Logger
}

func NewApplication(logger log.Logger) schedule.Schedule {
	return &Application{
		cron:   cron.New(cron.WithSeconds(), cron.WithChain(cron.Recover(cron.DefaultLogger))),
		logger: logger,
	}
}

func (a *Application) Register(jobs []schedule.Job) {
	for _, job := range jobs {
		_, err := a.cron.AddFunc(job.GetTiming(), job.GetCallback())

		if nil != err {
			a.logger.Print(err.Error())
		}
	}
}

func (a *Application) Run() {
	a.logger.Info("Scheduler runs with [Second | Minute | Hour | Dom | Month | Dow | Descriptor] parser.")
	a.cron.Run()
}
