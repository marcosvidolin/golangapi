package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"questionsandanswers/repository"
	r "questionsandanswers/router"

	"github.com/gorilla/mux"
)

func HttpHandler() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/questions/{id}", r.GetQuestionById).Methods("GET")
	router.HandleFunc("/questions", r.GetQuestionByAuthor).Methods("GET").Queries("author", "{author}")
	router.HandleFunc("/questions", r.GetAllQuestions).Methods("GET")
	router.HandleFunc("/questions", r.CreateQuestion).Methods("POST")
	router.HandleFunc("/questions/{id}", r.UpdateQuestion).Methods("PUT")
	router.HandleFunc("/questions/{id}", r.DeleteQuestion).Methods("DELETE")

	router.HandleFunc("/questions/{id}/answers", r.CreateAnswer).Methods("POST")
	router.HandleFunc("/questions/{id}/answers", r.UpdateAnswer).Methods("PUT")

	return router
}

func main() {

	dbclient, err := repository.GetMontoDbClient()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dbclient.Connect(context.Background()))

	router := HttpHandler()
	log.Fatal(http.ListenAndServe(":8080", router))
}
