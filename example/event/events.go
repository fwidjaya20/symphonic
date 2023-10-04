package event

import "time"

type PostCreated struct {
	Id        int64     `json:"id"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

func (pc PostCreated) Signature() string {
	return "Post.Created"
}

func (pc PostCreated) GetPayload() any {
	return pc
}
