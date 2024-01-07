package event

import (
	"context"

	ContractConfig "github.com/fwidjaya20/symphonic/contracts/config"
	ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"
	ContractLog "github.com/fwidjaya20/symphonic/contracts/log"
)

type Application struct {
	config    ContractConfig.Config
	driver    ContractEvent.QueueDriver
	listeners []ContractEvent.Listener
	logger    ContractLog.Logger
}

func (a *Application) Flush() error {
	return a.driver.Flush()
}

func (a *Application) Job(job ContractEvent.Job) ContractEvent.Bus {
	return NewEventBus(a.config, job, a.listeners, a.logger)
}

func (a *Application) Listeners() []ContractEvent.Listener {
	return a.listeners
}

func (a *Application) Register(listeners []ContractEvent.Listener) {
	a.listeners = listeners
}

func (a *Application) Run(config ContractEvent.RunEvent) error {
	a.driver = GetQueueDriver(config.Connection, &ContractEvent.DriverArgs{
		Config:        a.config,
		ConsumerGroup: config.ConsumerGroup,
		InitialOffset: config.Offset,
		Job:           nil,
		Listeners:     a.listeners,
		Logger:        a.logger,
		Topic:         config.Topic,
	})

	return a.driver.Subscribe(context.Background())
}

func NewApplication(config ContractConfig.Config, logger ContractLog.Logger) ContractEvent.Event {
	return &Application{
		config:    config,
		driver:    nil,
		listeners: make([]ContractEvent.Listener, 0),
		logger:    logger,
	}
}
