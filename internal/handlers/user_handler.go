package handlers

import (
	"fmt"
	"net/http"

	"github.com/WagaoCarvalho/quicknote/internal/repositories"
)

type userHandler struct {
	//	userRepository repositories.UserRepository
}

func NewUserHandler(userRepository repositories.UserRepository) *userHandler {
	return &userHandler{
		//userRepository: userRepository,
	}
}

func (uh *userHandler) SignupForm(w http.ResponseWriter, r *http.Request) error {
	render(w, http.StatusOK, "user_signup.html", nil)
	return nil
}

func (uh *userHandler) Signup(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erro ao processar o formulário", http.StatusBadRequest)
		return nil
	}

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	userRequest := &UserRequest{
		Email:    email,    // Mantém o valor digitado
		Password: password, // Mantém o valor digitado
	}

	if email == "" {
		userRequest.AddFieldError("email", "O e-mail é obrigatório")
	} else if !userRequest.IsEmailValid(email) {
		userRequest.AddFieldError("email", "O e-mail informado não é válido")
	}

	if password == "" {
		userRequest.AddFieldError("password", "A senha é obrigatória")
	} else if len(password) < 6 {
		userRequest.AddFieldError("password", "A senha deve ter pelo menos 6 caracteres")
	}

	// Se houver erros, renderiza novamente o formulário com os valores preenchidos
	if !userRequest.Valid() {
		render(w, http.StatusUnprocessableEntity, "user_signup.html", userRequest)
		return nil
	}

	fmt.Println(email, password)

	render(w, http.StatusOK, "user_signup.html", nil)
	return nil
}
