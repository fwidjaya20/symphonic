package event

type Event interface {
	Flush() error
	Job(job Job) Bus
	Listeners() []Listener
	Register(listeners []Listener)
	Run(config RunEvent) error
}

type Bus interface {
	OnConnection(connection string) Bus
	Publish() error
}

type RunEvent struct {
	Connection    string
	ConsumerGroup string
	Job           Job
	Offset        Offset
}
