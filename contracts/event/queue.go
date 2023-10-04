package event

import (
	"context"

	ContractConfig "github.com/fwidjaya20/symphonic/contracts/config"
	ContractLog "github.com/fwidjaya20/symphonic/contracts/log"
)

type DriverArgs struct {
	Config    ContractConfig.Config
	Job       Job
	Listeners []Listener
	Logger    ContractLog.Logger
	QueueName string
}

type QueueDriver interface {
	Driver() string
	Publish() error
	Subscribe(c context.Context) error
	Flush() error
}
