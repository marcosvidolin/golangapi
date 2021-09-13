package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Body      string             `json:"body,omitempty" bson:"body,omitempty"`
	Answer    Answer             `json:"answer" bson:"answer"`
	Author    User               `json:"author" bson:"author"`
	CreatedAt time.Time          `json:"-" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"-" bson:"updated_at"`
}

func NewQuestion() *Question {
	q := Question{}
	q.ID = primitive.NewObjectID()
	q.CreatedAt = time.Now()
	return &q
}

func (q *Question) AddAnswer(answer *Answer) error {

	if q.Answer != (Answer{}) {
		return ErrorQuestionAnswered
	}

	answer.CreatedAt = time.Now()

	q.Answer = *answer

	return nil
}

func (q *Question) UpdateAnswer(answer *Answer) error {

	if q.Answer == (Answer{}) {
		return ErrorNoAnswerToUpdate
	}

	if q.Author.Username != answer.Author.Username {
		return ErrorUnauthorizedUser
	}

	q.Answer.Body = answer.Body
	q.Answer.UpdatedAt = time.Now()

	return nil
}
