package transport

import "questionsandanswers/domain"

type FindQuestionByIdRequest struct {
	ID string `json:"id"`
}

type FindQuestionByIdResponse struct {
	Question domain.Question `json:"question"`
}
