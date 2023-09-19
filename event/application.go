package event

import (
	"github.com/fwidjaya20/symphonic/contracts/event"
)

type Application struct {
	events event.Collection
}

func NewApplication() event.Event {
	return &Application{
		events: make(event.Collection),
	}
}

func (a *Application) Collection() event.Collection {
	return a.events
}

func (a *Application) Job(job event.Job) event.Bus {
	return NewEventBus(job, a.events[job.Signature()])
}

func (a *Application) Register(events event.Collection) {
	a.events = events
}
