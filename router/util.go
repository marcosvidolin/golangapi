package router

import (
	"questionsandanswers/domain"
)

func GetCurrentUser() domain.User {
	return *domain.NewUser()
}
