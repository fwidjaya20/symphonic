package event

type Listener interface {
	Event() string
	Handle(job Job) error
}
