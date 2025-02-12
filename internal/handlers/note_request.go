package handlers

import (
	"fmt"

	"github.com/WagaoCarvalho/quicknote/internal/models"
	"github.com/WagaoCarvalho/quicknote/internal/validations"
)

// NoteRequest representa os dados necessários para a renderização da página de edição de notas
type NoteRequest struct {
	Id      int
	Title   string
	Content string
	Color   string
	Colors  []string
	validations.FormValidator
}

// newNoteRequest inicializa uma NoteRequest com base em uma nota existente
func (req *NoteRequest) newNoteRequest(note *models.Note) {
	// Adiciona cores pré-definidas
	for i := 1; i < 10; i++ {
		req.Colors = append(req.Colors, fmt.Sprintf("color%d", i))
	}

	// Preenche os campos da nota, se o ID for válido
	if note != nil && note.Id != 0 {
		req.Id = note.Id
		req.Title = note.Title
		req.Color = note.Color
		req.Content = note.Content
	} else {
		// Define valores padrão caso a nota seja nova
		req.Color = "color3"
	}
}
