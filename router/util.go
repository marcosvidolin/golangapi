package router

import (
	"questionsandanswers/domain"
)

func GetCurrentUser() domain.User {
	return domain.User{
		Username: "marcosvidolin",
		Email:    "mvidolin@xpto.com",
	}
}
