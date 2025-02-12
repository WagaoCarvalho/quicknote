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

func (req *UserRequest) newUserRequest(email, password string) {

	req.Email = email
	req.Password = password

}
