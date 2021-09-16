package transport

import "questionsandanswers/domain"

type FindAllQuestionsRequest struct{}

type FindAllQuestionsResponse struct {
	Questions []domain.Question `json:"questions"`
}

type FindQuestionByIdRequest struct {
	ID string `json:"id"`
}

type FindQuestionByIdResponse struct {
	Question domain.Question `json:"question"`
}

type FindQuestionsByAuthorRequest struct {
	Author string `json:"author"`
}

type FindQuestionsByAuthorResponse struct {
	Questions []domain.Question `json:"questions"`
}

type CreateQuestionRequest struct {
	Body string `json:"body"`
}

type CreateQuestionResponse struct {
	Question domain.Question `json:"questionn"`
}
