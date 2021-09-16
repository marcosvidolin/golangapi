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

// Creates a new Service
func NewService(rep repository.Repository) Service {
	return &service{
		repository: rep,
	}
}

// Creates a domain.question
// returns domainQuestion when there is no erros
// returns error when some validation error occurs
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

// Updates a domain.question
// returns domainQuestion when there is no erros
// returns error when some validation error occurs
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

// Gets a domain.Question by a given id
// returns a domain.Question found
// returns error when not found or something goes wrong
func (s *service) FindQuestionById(ctx context.Context, questionId string) (*domain.Question, error) {
	q, err := s.repository.FindQuestionById(ctx, questionId)

	if err != nil {
		return nil, err
	}

	return q, nil
}

// Gets all domain.Question
// returns all domain.Question found
// returns error when some thing goes wrong
func (s *service) FindAllQuestions(ctx context.Context) ([]domain.Question, error) {
	q, err := s.repository.FindAllQuestions(ctx)

	if err != nil {
		return nil, err
	}

	if len(*q) == 0 {
		empty := make([]domain.Question, 0)
		return empty, nil
	}

	return *q, nil
}

// Gets all domain.Question by a author (username)
// return a slice of domain.Questions found or an error when something goes wrong
func (s *service) FindQuestionsByAuthor(ctx context.Context, username string) ([]domain.Question, error) {
	q, err := s.repository.FindQuestionByAuthor(ctx, username)

	if err != nil {
		return nil, err
	}

	if len(*q) == 0 {
		empty := make([]domain.Question, 0)
		return empty, nil
	}

	return *q, nil
}

// Creates an domain.Answer for a given domain.Question
// returns domain.Question with the domain.Answer or error when something goes wrong
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

// Updates an domain.Answer for a given domain.Question
// returns domain.Question with the updated domain.Answer or error when something goes wrong
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

// Remove a domain.Question
// returns error when something goes wrong
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
