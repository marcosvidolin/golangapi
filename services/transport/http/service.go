package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"questionsandanswers/domain"
	"questionsandanswers/services/transport"

	kithttp "github.com/go-kit/kit/transport/http"
)

func NewService(endpoints transport.Endpoints) http.Handler {

	var (
		router = mux.NewRouter()
	)

	router.Methods("GET").Path("/questions").Handler(kithttp.NewServer(
		endpoints.GetAllQuestionsAndpoint,
		decodeGetAllQuestionsRequest,
		encodeResponse,
	))

	router.Methods("GET").Path("/questions/{id}").Handler(kithttp.NewServer(
		endpoints.GetQuestionByIdEndpoint,
		decodeGetQuestionByIdRequest,
		encodeResponse,
	))

	router.Methods("GET").Path("/auestions").Queries("author", "{author}").Handler(kithttp.NewServer(
		endpoints.GetQuestionsByAuthorEndpoint,
		decodeGetQuestionsByAuthorRequest,
		encodeResponse,
	))

	return router
}

func decodeGetAllQuestionsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return transport.FindAllQuestionsRequest{}, nil
}

func decodeGetQuestionByIdRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, domain.ErrorResourceNotFound
	}
	return transport.FindQuestionByIdRequest{ID: id}, nil
}

func decodeGetQuestionsByAuthorRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	author := r.FormValue("author")
	return transport.FindQuestionsByAuthorRequest{Author: author}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil errorr")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(getHttpStatusFromError(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func getHttpStatusFromError(err error) int {
	switch err {
	case domain.ErrorResourceNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
