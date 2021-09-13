package router

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"questionsandanswers/domain"
	"questionsandanswers/repository"
	"questionsandanswers/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var questionService service.Service

func init() {
	rep := repository.MongoDbRepository{}
	questionService = service.NewService(rep)
}

func GetQuestionById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	question, err := questionService.FindQuestionById(context.TODO(), id)

	if errors.Is(err, domain.ErrorResourceNotFound) {
		http.NotFound(w, r)
		return
	}

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if question == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(question)
}

func GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	questions, err := questionService.FindAllQuestions(context.TODO())

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(questions)
}

func GetQuestionByAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	author := r.FormValue("author")

	questions, err := questionService.FindQuestionsByAuthor(context.TODO(), author)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(questions)
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var q domain.Question

	err := json.NewDecoder(r.Body).Decode(&q)

	ctx := context.WithValue(context.Background(), "user", GetCurrentUser())

	questionService.CreateQuestion(ctx, &q)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(q)
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	var q *domain.Question
	err := json.NewDecoder(r.Body).Decode(&q)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	q.ID, err = primitive.ObjectIDFromHex(vars["id"])

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	ctx := context.WithValue(context.Background(), "user", GetCurrentUser())

	question, err := questionService.UpdateQuestion(ctx, q)

	if errors.Is(err, domain.ErrorResourceNotFound) {
		http.NotFound(w, r)
		return
	}

	if errors.Is(err, domain.ErrorUnauthorizedUser) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(question)
}

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	ctx := context.WithValue(context.Background(), "user", GetCurrentUser())

	err := questionService.DeleteQuestion(ctx, id)

	if errors.Is(err, domain.ErrorResourceNotFound) {
		http.NotFound(w, r)
		return
	}

	if errors.Is(err, domain.ErrorUnauthorizedUser) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}

func CreateAnswer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	var answer domain.Answer
	err := json.NewDecoder(r.Body).Decode(&answer)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	ctx := context.WithValue(context.Background(), "user", GetCurrentUser())

	var q *domain.Question
	q, err = questionService.CreateAnswer(ctx, id, &answer)

	if errors.Is(err, domain.ErrorQuestionAnswered) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if errors.Is(err, domain.ErrorUnauthorizedUser) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(q)
}

func UpdateAnswer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	var a domain.Answer
	err := json.NewDecoder(r.Body).Decode(&a)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	ctx := context.WithValue(context.Background(), "user", GetCurrentUser())

	q, err := questionService.UpdateAnswer(ctx, id, &a)

	if errors.Is(err, domain.ErrorNoAnswerToUpdate) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if errors.Is(err, domain.ErrorUnauthorizedUser) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(q)
}
