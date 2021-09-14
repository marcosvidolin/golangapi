package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type createQuestionRequest struct {
	ID primitive.ObjectID `json:"id"`
}

type createQuestionResponse struct{}

type updateQuestionRequest struct{}

type updateQuestionResponse struct{}

type findQuestionByIdRequest struct{}

type findQuestionByIdResponse struct{}

type findAllQuestionsRequest struct{}

type findAllQuestionsResponse struct{}

type findQuestionByAuthorRequest struct{}

type findQuestionByAuthorResponse struct{}

type createAnswerRequest struct{}

type createAnswerResponse struct{}

type updateAnswerRequest struct{}

type updateAnswerResponse struct{}

type deleteAnswerRequest struct{}

type deleteAnswerResponse struct{}
