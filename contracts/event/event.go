package event

type Event interface {
	Collection() Collection
	Job(job Job) Bus
	Register(events Collection)
	Run(config RunEvent) error
}

type Bus interface {
	OnConnection(connection string) Bus
	Publish() error
}

type Collection = map[string][]Listener

type RunEvent struct {
	Connection string
	Job        Job
}
