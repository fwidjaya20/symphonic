package event

type Job interface {
	GetPayload() any
	Signature() string
	Topic() string
}
