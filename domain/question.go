package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Body      string             `json:"body" bson:"body,omitempty"`
	Answer    Answer             `json:"answer" bson:"answer"`
	Author    User               `json:"author" bson:"author"`
	CreatedAt time.Time          `json:"-" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"-" bson:"updated_at"`
}

func NewQuestion() *Question {
	q := Question{}
	q.ID = primitive.NewObjectID()
	q.CreatedAt = time.Now()
	q.Author = *NewUser()
	return &q
}
