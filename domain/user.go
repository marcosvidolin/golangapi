package domain

type User struct {
	Username string `json:"username" bson:"username,omitempty"`
	Email    string `json:"-" bson:"email,emitempty"`
}

func NewUser() *User {
	author := User{
		Username: "marcosvidolin",
		Email:    "marcosvidolin@mail.com",
	}
	return &author
}
