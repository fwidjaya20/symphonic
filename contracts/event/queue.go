package event

import (
	"context"

	ContractConfig "github.com/fwidjaya20/symphonic/contracts/config"
	ContractLog "github.com/fwidjaya20/symphonic/contracts/log"
)

type DriverArgs struct {
	Config        ContractConfig.Config
	ConsumerGroup string
	InitialOffset Offset
	Job           Job
	Listeners     []Listener
	Logger        ContractLog.Logger
}

type QueueDriver interface {
	Driver() string
	Flush() error
	Publish() error
	Subscribe(ctx context.Context) error
}
