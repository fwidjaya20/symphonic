package event

type Event interface {
	Collection() Collection
	Job(job Job) Bus
	Register(events Collection)
}

type Bus interface {
	Publish() error
}

type Collection = map[string][]Listener
