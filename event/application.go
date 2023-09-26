package event

import (
	"context"

	"github.com/fwidjaya20/symphonic/contracts/config"
	"github.com/fwidjaya20/symphonic/contracts/event"
	"github.com/fwidjaya20/symphonic/contracts/log"
)

type Application struct {
	config config.Config
	events event.Collection
	logger log.Logger
}

func NewApplication(config config.Config, logger log.Logger) event.Event {
	return &Application{
		config: config,
		events: make(event.Collection),
		logger: logger,
	}
}

func (a *Application) Collection() event.Collection {
	return a.events
}

func (a *Application) Job(job event.Job) event.Bus {
	return NewEventBus(a.config, job, a.events[job.Signature()], a.logger)
}

func (a *Application) Register(events event.Collection) {
	a.events = events
}

func (a *Application) Run(config event.RunEvent) error {
	driver := GetQueueDriver(config.Connection, event.DriverArgs{
		Config:    a.config,
		Job:       config.Job,
		Listeners: a.Collection()[config.Job.Signature()],
		Logger:    a.logger,
	})

	return driver.Subscribe(context.Background())
}
