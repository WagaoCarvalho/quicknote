package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/WagaoCarvalho/quicknote/internal/models"
	"github.com/WagaoCarvalho/quicknote/internal/repositories"
)

type noteHandler struct {
	noteRepository repositories.NoteRepository
}

func NewNoteHandler(noteRepository repositories.NoteRepository) *noteHandler {
	return &noteHandler{noteRepository: noteRepository}
}

func (nh *noteHandler) NotesList(w http.ResponseWriter, r *http.Request) {
	// Verifica se o caminho da URL é a raiz "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Recupera a lista de notas do repositório
	notes, err := nh.noteRepository.ListNotes(r.Context())
	if err != nil {
		http.Error(w, "Aconteceu um erro ao acessar esta página", http.StatusInternalServerError)
		return
	}

	// Renderiza a página utilizando a função render
	render(w, http.StatusOK, "home.html", notes)
}

func (nh *noteHandler) NoteView(w http.ResponseWriter, r *http.Request) error {
	// Obtém o parâmetro "id" da URL
	idParam := r.PathValue("id")

	// Converte o ID para inteiro
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return err
	}

	// Recupera a nota pelo ID
	note, err := nh.noteRepository.GetNoteByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Erro ao recuperar a nota", http.StatusInternalServerError)
		return err
	}

	// Renderiza a página utilizando a função render
	render(w, http.StatusOK, "note_view.html", note)
	return nil
}

func (nh *noteHandler) NoteNew(w http.ResponseWriter, r *http.Request) error {
	// Criação de um objeto NoteColorRequest
	noteColorRequest := &NoteRequest{}
	noteColorRequest.newNoteRequest(nil) // Inicializa valores padrão, se necessário

	// Renderiza a página utilizando a função render
	render(w, http.StatusOK, "note_new.html", noteColorRequest)
	return nil
}

func (nh *noteHandler) NoteSave(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erro ao processar o formulário", http.StatusBadRequest)
		return nil
	}

	idParam := r.PostForm.Get("id")
	id, _ := strconv.Atoi(idParam)
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	color := r.PostForm.Get("color")

	var note *models.Note
	if id > 0 {
		note, err = nh.noteRepository.UpdateNote(r.Context(), id, title, content, color)
	} else {

		note, err = nh.noteRepository.CreateNote(r.Context(), title, content, color)
	}

	if err != nil {
		http.Error(w, "Erro ao criar a nota", http.StatusInternalServerError)
		return nil
	}

	http.Redirect(w, r, fmt.Sprintf("/note/%d", note.Id), http.StatusSeeOther)
	return nil
}

func (nh *noteHandler) NoteDelete(w http.ResponseWriter, r *http.Request) error {

	idParam := r.PathValue("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return nil
	}

	err = nh.noteRepository.DeleteNote(r.Context(), id)
	if err != nil {
		http.Error(w, "Erro ao deletar a nota", http.StatusInternalServerError)
		return nil
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}

// NoteEdit manipula a exibição da página de edição de notas
func (nh *noteHandler) NoteEdit(w http.ResponseWriter, r *http.Request) error {
	// Obtém o ID da nota a partir da URL
	idParam := r.PathValue("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return nil
	}

	// Recupera a nota pelo ID
	note, err := nh.noteRepository.GetNoteByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Erro ao encontrar a nota", http.StatusInternalServerError)
		return nil
	}

	// Inicializa o objeto NoteRequest com base na nota recuperada
	noteRequest := &NoteRequest{}
	noteRequest.newNoteRequest(note)

	// Renderiza a página de edição com os dados da nota
	render(w, http.StatusOK, "note_edit.html", noteRequest)
	return nil
}
