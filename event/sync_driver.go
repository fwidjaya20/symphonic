package event

import (
	"context"
	"encoding/json"
	"sync"

	ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"
	"github.com/sirupsen/logrus"
)

type SyncDriver struct {
	*ContractEvent.DriverArgs
}

func (d *SyncDriver) Driver() string {
	return DriverSync
}

func (d *SyncDriver) Publish() error {
	var errors []error

	var mu sync.Mutex

	var wg sync.WaitGroup

	d.Logger.Infof("'%s' has been published: %v", d.Job.Signature(), d.Job.GetPayload())

	payload, err := json.Marshal(d.Job.GetPayload())
	if err != nil {
		d.Logger.WithFields(logrus.Fields{
			logrus.ErrorKey: err,
		}).Errorf("Unable to marshal %v", d.Job.GetPayload())

		return err
	}

	for _, listener := range d.Listeners {
		wg.Add(1)

		go func(callback ContractEvent.Listener) {
			defer wg.Done()

			if err := callback.Handle(payload); err != nil {
				mu.Lock()
				defer mu.Unlock()

				errors = append(errors, err)
			}
		}(listener)
	}

	wg.Wait()

	return nil
}

func (d *SyncDriver) Subscribe(_ context.Context) error {
	d.Logger.Infof(
		"Running the Sync Driver explicitly is unnecessary and could potentially disrupt system operations.",
	)

	return nil
}

func (d *SyncDriver) Flush() error {
	return nil
}

func NewSyncDriver(args *ContractEvent.DriverArgs) ContractEvent.QueueDriver {
	return &SyncDriver{
		args,
	}
}
