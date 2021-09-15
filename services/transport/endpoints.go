package transport

import (
	"context"

	"questionsandanswers/services"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetQuestionByIdEndpoint endpoint.Endpoint
}

func MakeEndpoints(s services.Service) Endpoints {
	return Endpoints{
		GetQuestionByIdEndpoint: makeGetQuestionByIdEndpoint(s),
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
