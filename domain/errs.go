package domain

import "errors"

var ErrorResourceNotFound error
var ErrorUnauthorizedUser error
var ErrorQuestionAnswered error
var ErrorNoAnswerToUpdate error

func init() {
	ErrorResourceNotFound = errors.New("RESOURCE NOT FOUND")
	ErrorUnauthorizedUser = errors.New("UNAUTORIZED USER")
	ErrorQuestionAnswered = errors.New("QUESTION ALREADY ANSWERED")
	ErrorNoAnswerToUpdate = errors.New("NO ANSWER TO UPDATE")
}
