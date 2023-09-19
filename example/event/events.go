package event

import "time"

type PostCreated struct {
	Id        int64
	Author    string
	CreatedAt time.Time
}

func (pc PostCreated) Signature() string {
	return "Post.Created"
}

func (pc PostCreated) GetPayload() any {
	return pc
}
