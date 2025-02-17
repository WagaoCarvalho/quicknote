package handlers

import (
	"fmt"
	"net/http"

	"github.com/WagaoCarvalho/quicknote/repositories"
	"github.com/WagaoCarvalho/quicknote/utils"
)

type userHandler struct {
	userRepository repositories.UserRepository
	passwordHasher utils.BcryptHasher
}

func NewUserHandler(userRepository repositories.UserRepository) *userHandler {
	return &userHandler{
		userRepository: userRepository,
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

	userRequest := newUserRequest(email, password)

	if userRequest.Email == "" {
		userRequest.AddFieldError("email", "O e-mail é obrigatório")
	} else if !userRequest.IsEmailValid(userRequest.Email) {
		userRequest.AddFieldError("email", "O e-mail informado não é válido")
	}

	if userRequest.Password == "" {
		userRequest.AddFieldError("password", "A senha é obrigatória")
	} else if len(userRequest.Password) < 6 {
		userRequest.AddFieldError("password", "A senha deve ter pelo menos 6 caracteres")
	}

	if !userRequest.Valid() {
		render(w, http.StatusUnprocessableEntity, "user_signup.html", userRequest)
		return nil
	}

	//hash
	// Injetando a implementação real

	hashedPassword, err := uh.passwordHasher.HashPassword(userRequest.Password)
	if err != nil {
		http.Error(w, "Erro ao gerar hash da senha", http.StatusInternalServerError)
		return nil
	}
	fmt.Println(hashedPassword)

	user, err := uh.userRepository.CreateUser(r.Context(), userRequest.Email, hashedPassword)

	if err == repositories.ErrDuplicateEmail {
		userRequest.AddFieldError("email", "O email Já está em uso")
		render(w, http.StatusUnprocessableEntity, "user_signup.html", userRequest)
		return nil
	}

	if err != nil {
		http.Error(w, "Erro ao criar o usuário", http.StatusInternalServerError)
		return nil
	}

	fmt.Println("Usuário Criado com sucesso!! Id: ", user.Id)

	render(w, http.StatusOK, "user_signup_success.html", nil)
	return nil
}
