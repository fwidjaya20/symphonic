package events

import "time"

type PostCreated struct {
	ID        int64     `json:"id"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

func (pc PostCreated) GetPayload() any {
	return pc
}

func (pc PostCreated) Signature() string {
	return "Post.Created"
}

func (pc PostCreated) Topic() string {
	return "Post.Created"
}
