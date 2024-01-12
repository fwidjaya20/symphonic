package event

import (
	ContractConfig "github.com/fwidjaya20/symphonic/contracts/config"
	"github.com/fwidjaya20/symphonic/contracts/event"
	"github.com/fwidjaya20/symphonic/contracts/log"
)

type Bus struct {
	config     ContractConfig.Config
	connection string
	driver     event.QueueDriver
	isQueued   bool
	job        event.Job
	listeners  []event.Listener
	logger     log.Logger
}

func NewEventBus(
	config ContractConfig.Config,
	job event.Job,
	listeners []event.Listener,
	logger log.Logger,
) event.Bus {
	connection := config.GetString("queue.default", DriverSync)

	bus := Bus{
		config:     config,
		connection: connection,
		driver:     nil,
		isQueued:   false,
		job:        job,
		listeners:  listeners,
		logger:     logger,
	}

	bus.OnConnection(connection)

	return &bus
}

func (b *Bus) OnConnection(driver string) event.Bus {
	b.provideDriver(driver)

	if driver != DriverSync {
		b.isQueued = true
	}

	b.connection = driver

	return b
}

func (b *Bus) Publish() error {
	return b.driver.Publish()
}

func (b *Bus) provideDriver(driver string) {
	b.driver = GetQueueDriver(driver, &event.DriverArgs{
		Config:        b.config,
		ConsumerGroup: "",
		InitialOffset: event.OffsetOldest,
		Job:           b.job,
		Listeners:     b.listeners,
		Logger:        b.logger,
	})
}
