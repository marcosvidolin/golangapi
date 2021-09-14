package service

import (
	"context"
	"questionsandanswers/domain"
	"questionsandanswers/repository"
	"strings"

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

func (s *service) CreateQuestion(ctx context.Context, question *domain.Question) (*domain.Question, error) {

	if len(strings.TrimSpace(question.Body)) == 0 {
		return nil, domain.ErrorContentRequired
	}

	question.Author = ctx.Value("user").(domain.User)
	question.CreatedAt = time.Now()

	entity, err := s.repository.CreateQuestion(ctx, question)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *service) UpdateQuestion(ctx context.Context, question *domain.Question) (*domain.Question, error) {
	q, err := s.FindQuestionById(ctx, question.ID.Hex())

	if err != nil {
		return nil, err
	}

	if q == nil {
		return nil, nil
	}

	if len(strings.TrimSpace(question.Body)) == 0 {
		return nil, domain.ErrorContentRequired
	}

	q.UpdatedAt = time.Now()
	q.Body = question.Body
	q.Answer = question.Answer

	author := ctx.Value("user").(domain.User)
	if q.Author.Username != author.Username {
		return nil, domain.ErrorUnauthorizedUser
	}

	qst, err := s.repository.UpdateQuestion(ctx, q)

	if err != nil {
		return nil, err
	}

	return qst, nil
}

func (s *service) FindQuestionById(ctx context.Context, questionId string) (*domain.Question, error) {
	q, err := s.repository.FindQuestionById(ctx, questionId)

	if err != nil {
		return nil, err
	}

	return q, nil
}

func (s *service) FindAllQuestions(ctx context.Context) (*[]domain.Question, error) {
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

func (s *service) FindQuestionsByAuthor(ctx context.Context, username string) (*[]domain.Question, error) {
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

func (s *service) CreateAnswer(ctx context.Context, questionId string, answer *domain.Answer) (*domain.Question, error) {
	q, err := s.repository.FindQuestionById(ctx, questionId)

	if err != nil {
		return nil, err
	}

	author := ctx.Value("user").(domain.User)
	answer.Author = author

	if err = q.AddAnswer(answer); err != nil {
		return nil, err
	}

	question, err := s.UpdateQuestion(ctx, q)

	if err != nil {
		return nil, err
	}

	return question, nil
}

func (s *service) UpdateAnswer(ctx context.Context, questionId string, answer *domain.Answer) (*domain.Question, error) {
	q, err := s.repository.FindQuestionById(ctx, questionId)

	if err != nil {
		return nil, err
	}

	author := ctx.Value("user").(domain.User)
	answer.Author = author

	if err = q.UpdateAnswer(answer); err != nil {
		return nil, err
	}

	entity, err := s.repository.UpdateQuestion(ctx, q)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *service) DeleteQuestion(ctx context.Context, id string) error {
	q, err := s.repository.FindQuestionById(ctx, id)

	if err != nil {
		return err
	}

	author := ctx.Value("user").(domain.User)

	if q.Author.Username != author.Username {
		return domain.ErrorUnauthorizedUser
	}

	return s.repository.DeleteQuestion(ctx, id)
}
