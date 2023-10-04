package event

import (
	"context"
	"sync"

	ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"
)

type SyncDriver struct {
	ContractEvent.DriverArgs
}

func NewSyncDriver(args ContractEvent.DriverArgs) ContractEvent.QueueDriver {
	return &SyncDriver{
		args,
	}
}

func (d *SyncDriver) Driver() string {
	return DriverSync
}

func (d *SyncDriver) Publish() error {
	var errors []error
	var mu sync.Mutex
	var wg sync.WaitGroup

	d.Logger.Infof("'%s' has been published: %v", d.Job.Signature(), d.Job.GetPayload())

	wg.Add(len(d.Listeners))

	for _, listener := range d.Listeners {
		go func(callback ContractEvent.Listener) {
			defer wg.Done()
			if err := callback.Handle(d.Job); nil != err {
				mu.Lock()
				defer mu.Unlock()
				errors = append(errors, err)
			}
		}(listener)
	}

	wg.Wait()

	return nil
}

func (d *SyncDriver) Subscribe(c context.Context) error {
	d.Logger.Infof("Running the Sync Driver explicitly is unnecessary and could potentially disrupt system operations.")
	return nil
}

func (d *SyncDriver) Flush() error {
	return nil
}
