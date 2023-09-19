package event

import (
	"fmt"
	"sync"

	"github.com/fwidjaya20/symphonic/contracts/event"
)

type Bus struct {
	event     event.Job
	listeners []event.Listener
	locker    sync.Mutex
}

func NewEventBus(event event.Job, listeners []event.Listener) event.Bus {
	return &Bus{
		event:     event,
		listeners: listeners,
	}
}

func (t *Bus) Publish() error {
	t.locker.Lock()
	defer t.locker.Unlock()

	if len(t.listeners) < 1 {
		return fmt.Errorf("event '%s' doesn't bind any listeners", t.event.Signature())
	}

	var errors []error
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(len(t.listeners))

	for _, listener := range t.listeners {
		go func(callback event.Listener) {
			defer wg.Done()
			if err := callback.Handle(t.event); nil != err {
				mu.Lock()
				defer mu.Unlock()
				errors = append(errors, err)
			}
		}(listener)
	}

	wg.Wait()

	return nil
}
