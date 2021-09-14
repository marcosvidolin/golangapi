package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewQuestion(t *testing.T) {
	question := NewQuestion()

	assert.NotEmpty(t, question.ID)
	assert.NotEmpty(t, question.CreatedAt)
}

func TestCreateAnswer(t *testing.T) {
	author := User{Username: "marcosvidolin"}

	question := NewQuestion()
	question.Author = author

	answer := NewAnswer()
	answer.Body = "The answer for this question is..."
	answer.Author = author

	err := question.AddAnswer(answer)

	assert.Empty(t, err)
	assert.Equal(t, author.Username, question.Answer.Author.Username)
	assert.NotNil(t, answer.Body)
	assert.Equal(t, answer.Body, question.Answer.Body)
	assert.NotNil(t, question.Answer.UpdatedAt)
}

func TestCreateAnswerQuestionAlreadyAnswered(t *testing.T) {
	author := User{Username: "marcosvidolin"}

	question := NewQuestion()
	question.Answer = *NewAnswer()
	question.Author = author

	answer := NewAnswer()
	answer.Body = "The answer for this question is..."
	answer.Author = author

	err := question.AddAnswer(answer)

	assert.NotNil(t, err)
	assert.Equal(t, ErrorQuestionAnswered, err)
}

func TestUpdateAnswer(t *testing.T) {
	author := User{Username: "marcosvidolin"}

	question := NewQuestion()
	question.Answer = *NewAnswer()
	question.Author = author

	answer := NewAnswer()
	answer.Body = "The answer for this question is..."
	answer.Author = author

	err := question.UpdateAnswer(answer)

	assert.Empty(t, err)
	assert.NotEmpty(t, question.Answer.Body)
	assert.Equal(t, "The answer for this question is...", question.Answer.Body)
	assert.NotNil(t, question.Answer.CreatedAt)
	assert.Equal(t, question.Author.Username, answer.Author.Username)
}

func TestUpdateAnswerThereIsNoAnswerToUpdate(t *testing.T) {
	author := User{Username: "marcosvidolin"}
	question := NewQuestion()
	question.Author = author

	err := question.UpdateAnswer(&Answer{Author: author})

	assert.NotNil(t, err)
	assert.Equal(t, ErrorNoAnswerToUpdate, err)
}

func TestUpdateAnswerUserUnauthorized(t *testing.T) {
	question := NewQuestion()
	question.Author = User{Username: "marcosvidolin"}
	question.Answer = *NewAnswer()

	author := User{Username: "anonymous"}

	err := question.UpdateAnswer(&Answer{Author: author, Body: "Just do it..."})

	assert.NotNil(t, err)
	assert.Equal(t, ErrorUnauthorizedUser, err)
}
