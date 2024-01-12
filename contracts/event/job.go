package event

type Job interface {
	Event() string
	GetPayload() any
	Topic() string
}
