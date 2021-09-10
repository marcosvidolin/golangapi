package service

import (
	"context"
	"errors"
	"questionsandanswers/domain"
	"questionsandanswers/repository"

	"time"
)

type service struct {
	repository repository.Repository
}

func NewService(rep repository.Repository) Service {
	return &service{
		repository: rep,
	}
}

func (s service) CreateQuestion(ctx context.Context, question *domain.Question) (*domain.Question, error) {

	question.CreatedAt = time.Now()

	entity, err := s.repository.CreateQuestion(ctx, question)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (s service) UpdateQuestion(ctx context.Context, question domain.Question) (*domain.Question, error) {
	q, err := s.FindQuestionById(ctx, question.ID.Hex())

	if err != nil {
		return nil, err
	}

	if q == nil {
		return nil, nil
	}

	q.UpdatedAt = time.Now()
	q.Body = question.Body
	q.Answer = question.Answer

	qst, err := s.repository.UpdateQuestion(ctx, *q)

	if err != nil {
		return nil, err
	}

	return qst, nil
}

func (s service) FindQuestionById(ctx context.Context, questionId string) (*domain.Question, error) {
	q, err := s.repository.FindQuestionById(ctx, questionId)

	if err != nil {
		return nil, err
	}

	return q, nil
}

func (s service) FindAllQuestions(ctx context.Context) (*[]domain.Question, error) {
	q, err := s.repository.FindAllQuestions(ctx)

	if err != nil {
		return nil, err
	}

	if q == nil || len(*q) == 0 {
		empty := make([]domain.Question, 0)
		return &empty, nil
	}

	return q, nil
}

func (s service) FindQuestionsByAuthor(ctx context.Context, username string) (*[]domain.Question, error) {
	q, err := s.repository.FindQuestionByAuthor(ctx, username)

	if err != nil {
		return nil, err
	}

	if q == nil || len(*q) == 0 {
		empty := make([]domain.Question, 0)
		return &empty, nil
	}

	return q, nil
}

func (s service) CreateAnswer(ctx context.Context, questionId string, answer *domain.Answer) (*domain.Question, error) {
	q, err := s.repository.FindQuestionById(ctx, questionId)

	if err != nil {
		return nil, err
	}

	if q.Answer != (domain.Answer{}) {
		return nil, domain.ErrorQuestionAnswered
	}

	answer.CreatedAt = time.Now()
	q.Answer = *answer

	question, err := s.UpdateQuestion(ctx, *q)

	if err != nil {
		return nil, err
	}

	return question, nil
}

func (s service) UpdateAnswer(ctx context.Context, questionId string, answer *domain.Answer) (*domain.Question, error) {
	q, err := s.repository.FindQuestionById(ctx, questionId)

	if err != nil {
		return nil, err
	}

	if q.Answer == (domain.Answer{}) {
		return nil, domain.ErrorNoAnswerToUpdate
	}

	if q.Author.Username != answer.Author.Username {
		return nil, errors.New("UNAUTHORIZED USER")
	}

	answer.UpdatedAt = time.Now()
	q.Answer.Body = answer.Body

	entity, err := s.repository.UpdateQuestion(ctx, *q)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (s service) DeleteQuestion(ctx context.Context, id string) error {
	return s.repository.DeleteQuestion(ctx, id)
}
