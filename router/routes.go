package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HttpHandler() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/questions/{id}", GetQuestionById).Methods("GET")
	router.HandleFunc("/questions", GetQuestionByAuthor).Methods("GET").Queries("author", "{author}")
	router.HandleFunc("/questions", GetAllQuestions).Methods("GET")
	router.HandleFunc("/questions", CreateQuestion).Methods("POST")
	router.HandleFunc("/questions/{id}", UpdateQuestion).Methods("PUT")
	router.HandleFunc("/questions/{id}", DeleteQuestion).Methods("DELETE")

	router.HandleFunc("/questions/{id}/answers", CreateAnswer).Methods("POST")
	router.HandleFunc("/questions/{id}/answers", UpdateAnswer).Methods("PUT")

	return router
}
