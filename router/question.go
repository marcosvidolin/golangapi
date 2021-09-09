package router

import (
	"context"
	"encoding/json"
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
	vars := mux.Vars(r)
	id := vars["id"]

	question, _ := questionService.FindQuestionById(context.TODO(), id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(question)
}

func GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions, _ := questionService.FindAllQuestions(context.TODO())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func GetQuestionByAuthor(w http.ResponseWriter, r *http.Request) {
	author := r.FormValue("author")

	questions, err := questionService.FindQuestionsByAuthor(context.TODO(), author)

	if err != nil {

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {

	var q domain.Question

	err := json.NewDecoder(r.Body).Decode(&q)

	questionService.CreateQuestion(context.TODO(), q)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(q)
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var q domain.Question
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

	question, err := questionService.UpdateQuestion(context.TODO(), q)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(question)
}

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := questionService.DeleteQuestion(context.TODO(), id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("") // TODO: empty
}

func CreateAnswer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	question, err := questionService.FindQuestionById(context.TODO(), id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var a domain.Answer
	err = json.NewDecoder(r.Body).Decode(&a)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	answer := *domain.NewAnswer()
	answer.Body = a.Body
	question.Answer = answer

	q, err := questionService.UpdateQuestion(context.TODO(), *question)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(q)
}

func UpdateAnswer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	question, err := questionService.FindQuestionById(context.TODO(), id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var a domain.Answer
	err = json.NewDecoder(r.Body).Decode(&a)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	question.Answer.Body = a.Body
	// TODO: packagessar update para answer e atualizar a data update

	q, err := questionService.UpdateQuestion(context.TODO(), *question)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(q)
}