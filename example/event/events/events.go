package events

import "time"

const (
	PostCreatedEvent = "created"
	PostCreatedTopic = "post.created"
)

type PostCreated struct {
	ID        int64     `json:"id"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

func (pc PostCreated) GetPayload() any {
	return pc
}

func (pc PostCreated) Event() string {
	return PostCreatedEvent
}

func (pc PostCreated) Topic() string {
	return PostCreatedTopic
}
