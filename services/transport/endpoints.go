package transport

import (
	"context"

	"questionsandanswers/domain"
	"questionsandanswers/services"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetAllQuestionsAndpoint      endpoint.Endpoint
	GetQuestionByIdEndpoint      endpoint.Endpoint
	GetQuestionsByAuthorEndpoint endpoint.Endpoint
	CreateQuestionEndpoint       endpoint.Endpoint
}

func MakeEndpoints(s services.Service) Endpoints {
	return Endpoints{
		GetAllQuestionsAndpoint:      makeGetAllQuestionsEndpoint(s),
		GetQuestionByIdEndpoint:      makeGetQuestionByIdEndpoint(s),
		GetQuestionsByAuthorEndpoint: makeGetQuestionsByAuthorEndpoint(s),
		CreateQuestionEndpoint:       makeCreateQuestionEndpoint(s),
	}
}

func makeGetAllQuestionsEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllQuestionsRequest)

		q, err := s.FindAllQuestions(ctx)

		if err != nil {
			return FindAllQuestionsResponse{*q}, err
		}

		return FindAllQuestionsResponse{*q}, nil
	}
}

func makeGetQuestionByIdEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FindQuestionByIdRequest)

		q, err := s.FindQuestionById(ctx, req.ID)

		if err != nil {
			return FindQuestionByIdResponse{*q}, err
		}

		return FindQuestionByIdResponse{*q}, nil
	}
}

func makeGetQuestionsByAuthorEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FindQuestionsByAuthorRequest)

		q, err := s.FindQuestionsByAuthor(ctx, req.Author)

		if err != nil {
			return FindQuestionsByAuthorResponse{*q}, err
		}

		return FindQuestionsByAuthorResponse{*q}, nil
	}
}

func makeCreateQuestionEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateQuestionRequest)

		question := domain.Question{
			Body: req.Body,
		}

		q, err := s.CreateQuestion(ctx, &question)

		if err != nil {
			return CreateQuestionResponse{*q}, err
		}

		return CreateQuestionResponse{*q}, err
	}
}
