package domain

import "errors"

var ErrorContentRequired error
var ErrorResourceNotFound error
var ErrorUnauthorizedUser error
var ErrorQuestionAnswered error
var ErrorNoAnswerToUpdate error

func init() {
	ErrorContentRequired = errors.New("CONTENT REQUIRED")
	ErrorResourceNotFound = errors.New("RESOURCE NOT FOUND")
	ErrorUnauthorizedUser = errors.New("UNAUTORIZED USER")
	ErrorQuestionAnswered = errors.New("QUESTION ALREADY ANSWERED")
	ErrorNoAnswerToUpdate = errors.New("NO ANSWER TO UPDATE")
}
