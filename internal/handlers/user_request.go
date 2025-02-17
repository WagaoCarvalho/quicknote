package handlers

import "github.com/WagaoCarvalho/quicknote/internal/validations"

type UserRequest struct {
	Email    string
	Password string
	Content  string
	Color    string
	Colors   []string
	validations.FormValidator
}

func newUserRequest(email, password string) *UserRequest {
	return &UserRequest{
		Email:    email,
		Password: password,
	}
}
