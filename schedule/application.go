package schedule

import (
	"time"

	"github.com/fwidjaya20/symphonic/contracts/log"
	"github.com/fwidjaya20/symphonic/contracts/schedule"
	"github.com/robfig/cron/v3"
)

type Application struct {
	cron   *cron.Cron
	logger log.Logger
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
	a.logger.Info("Scheduler runs with [SecondOptional | Minute | Hour | Dom | Month | Dow | Descriptor] parser.")
	a.cron.Start()
}

func (a *Application) Stop() {
	a.cron.Stop()
}

func NewApplication(logger log.Logger) schedule.Schedule {
	return &Application{
		cron: cron.New(
			cron.WithParser(
				cron.NewParser(cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dom|cron.Descriptor),
			),
			cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)),
			cron.WithChain(cron.Recover(cron.DefaultLogger)),
			cron.WithLocation(time.UTC),
		),
		logger: logger,
	}
}
