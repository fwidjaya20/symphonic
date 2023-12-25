package event

import (
	"context"

	"github.com/fwidjaya20/symphonic/contracts/config"
	"github.com/fwidjaya20/symphonic/contracts/event"
	"github.com/fwidjaya20/symphonic/contracts/log"
)

type Application struct {
	config    config.Config
	driver    event.QueueDriver
	listeners event.Collection
	logger    log.Logger
}

func (a *Application) Collection() event.Collection {
	return a.listeners
}

func (a *Application) Flush() error {
	return a.driver.Flush()
}

func (a *Application) Job(job event.Job) event.Bus {
	return NewEventBus(a.config, job, a.listeners[job.Topic()], a.logger)
}

func (a *Application) Register(listeners event.Collection) {
	a.listeners = listeners
}

func (a *Application) Run(config event.RunEvent) error {
	a.driver = GetQueueDriver(config.Connection, event.DriverArgs{
		Config:        a.config,
		ConsumerGroup: config.ConsumerGroup,
		InitialOffset: config.Offset,
		Job:           config.Job,
		Listeners:     a.Collection()[config.Job.Signature()],
		Logger:        a.logger,
	})

	return a.driver.Subscribe(context.Background())
}

func NewApplication(config config.Config, logger log.Logger) event.Event {
	return &Application{
		config:    config,
		listeners: make(event.Collection),
		logger:    logger,
	}
}
