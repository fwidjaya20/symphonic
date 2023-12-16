package event

import (
	"sync"

	"github.com/fwidjaya20/symphonic/contracts/config"
	"github.com/fwidjaya20/symphonic/contracts/event"
	"github.com/fwidjaya20/symphonic/contracts/log"
)

type Bus struct {
	connection string
	config     config.Config
	driver     event.QueueDriver
	event      event.Job
	isQueued   bool
	listeners  []event.Listener
	locker     sync.Mutex
	logger     log.Logger
	queueName  string
}

func NewEventBus(config config.Config, event event.Job, listeners []event.Listener, logger log.Logger) event.Bus {
	connection := config.GetString("queue.default", DriverSync)

	bus := &Bus{
		connection: connection,
		config:     config,
		event:      event,
		listeners:  listeners,
		logger:     logger,
	}

	bus.determineDriver(connection)
	bus.determineQueue(connection)

	return bus
}

func (b *Bus) OnConnection(driver string) event.Bus {
	b.determineDriver(driver)
	b.determineQueue(driver)
	b.connection = driver
	return b
}

func (b *Bus) OnQueue(queueName string) event.Bus {
	b.queueName = queueName
	return b
}

func (b *Bus) Publish() error {
	b.locker.Lock()
	defer b.locker.Unlock()

	if len(b.listeners) == 0 {
		b.logger.Infof("event %s doesn't bind any listeners", b.event.Signature())
	}

	return b.driver.Publish()
}

func (b *Bus) determineQueue(driver string) {
	switch driver {
	case DriverSync:
		b.isQueued = false
	default:
		b.isQueued = true
	}
}

func (b *Bus) determineDriver(driver string) {
	b.driver = GetQueueDriver(driver, event.DriverArgs{
		Config:    b.config,
		Job:       b.event,
		Listeners: b.listeners,
		Logger:    b.logger,
		QueueName: b.queueName,
	})
}
