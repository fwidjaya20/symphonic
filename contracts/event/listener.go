package event

type Listener interface {
	Handle(messages []byte) error
	Signature() string
}
