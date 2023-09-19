package event

type Listener interface {
	Handle(event Job) error
	Signature() string
}
