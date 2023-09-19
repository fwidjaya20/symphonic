package event

type Job interface {
	Signature() string
	GetPayload() any
}
