package domain

import "time"

type Answer struct {
	Body      string    `json:"body" bson:"body,omitempty"`
	Author    User      `json:"author" bson:"author,omitempty"`
	CreatedAt time.Time `json:"-" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"-" bson:"updated_at"`
}

func NewAnswer() *Answer {
	a := Answer{}
	a.CreatedAt = time.Now()
	return &a
}
