package services

import (
	"context"

	"questionsandanswers/domain"
)

type Service interface {
	CreateQuestion(ctx context.Context, question *domain.Question) (*domain.Question, error)
	UpdateQuestion(ctx context.Context, question *domain.Question) (*domain.Question, error)
	DeleteQuestion(ctx context.Context, id string) error
	FindQuestionById(ctx context.Context, id string) (*domain.Question, error)
	FindAllQuestions(ctx context.Context) (*[]domain.Question, error)
	FindQuestionsByAuthor(ctx context.Context, username string) (*[]domain.Question, error)
	CreateAnswer(ctx context.Context, questionId string, answer *domain.Answer) (*domain.Question, error)
	UpdateAnswer(ctx context.Context, questionId string, answer *domain.Answer) (*domain.Question, error)
}
