package schedule

import (
	"log"

	"github.com/fwidjaya20/symphonic/contracts/schedule"
	"github.com/robfig/cron/v3"
)

type Application struct {
	cron *cron.Cron
}

func NewApplication() schedule.Schedule {
	return &Application{
		cron: cron.New(cron.WithSeconds(), cron.WithChain(cron.Recover(cron.DefaultLogger))),
	}
}

func (a *Application) Register(jobs []schedule.Job) {
	for _, job := range jobs {
		_, err := a.cron.AddFunc(job.GetTiming(), job.GetCallback())

		if nil != err {
			log.Println(err.Error())
		}
	}
}

func (a *Application) Run() {
	a.cron.Run()
}
