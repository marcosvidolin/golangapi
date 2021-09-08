package service

import (
	"context"
	"fmt"
	"log"
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

func (s service) CreateQuestion(ctx context.Context, question domain.Question) (*domain.Question, error) {
	q := domain.NewQuestion()
	q.Body = question.Body
	entity, err := s.repository.CreateQuestion(ctx, *q)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (s service) UpdateQuestion(ctx context.Context, question domain.Question) (*domain.Question, error) {
	q, err := s.FindQuestionById(ctx, question.ID.Hex())

	fmt.Println(question.ID.Hex())

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	q.UpdatedAt = time.Now()
	q.Body = question.Body
	q.Answer = question.Answer

	qst, err := s.repository.UpdateQuestion(ctx, *q)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return qst, nil
}

func (s service) FindQuestionById(ctx context.Context, questionId string) (*domain.Question, error) {
	q, err := s.repository.FindQuestionById(ctx, questionId)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return q, nil
}

func (s service) FindAllQuestions(ctx context.Context) (*[]domain.Question, error) {
	q, err := s.repository.FindAllQuestions(ctx)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return q, nil
}

func (s service) FindQuestionsByAuthor(ctx context.Context, username string) (*[]domain.Question, error) {
	q, err := s.repository.FindQuestionByAuthor(ctx, username)

	if err != nil {
		return nil, err
	}

	return q, nil
}

func (s service) CreateAnswer(ctx context.Context, questionId string, answer domain.Answer) (*domain.Question, error) {
	q, err := s.repository.FindQuestionById(ctx, questionId)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	answer.CreatedAt = time.Now()
	q.Answer = answer

	question, err := s.UpdateQuestion(ctx, *q)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return question, nil
}

func (s service) UpdateAnswer(ctx context.Context, questionId string, answer domain.Answer) (*domain.Question, error) {
	q, err := s.repository.FindQuestionById(ctx, questionId)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	answer.UpdatedAt = time.Now()
	q.Answer = answer // TODO: atualizar somente o body

	entity, err := s.repository.UpdateQuestion(ctx, *q)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return entity, nil
}

func (s service) DeleteQuestion(ctx context.Context, id string) error {
	err := s.repository.DeleteQuestion(ctx, id)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
