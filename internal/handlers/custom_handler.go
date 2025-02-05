package handlers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/WagaoCarvalho/quicknote/internal/handlers/apperror"
)

var ErrNotFound = apperror.WithStatus(errors.New("não encontrado"), http.StatusNotFound)

type HandlerWithError func(w http.ResponseWriter, r *http.Request) error

func (f HandlerWithError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		var statusErr apperror.StatusError
		if errors.As(err, &statusErr) {
			if statusErr.StatusCode() == http.StatusNotFound {
				filesPages := []string{
					"views/templates/base.html",
					"views/templates/pages/note_view.html",
				}
				t, err := template.ParseFiles(filesPages...)
				if err != nil {
					http.Error(w, err.Error(), statusErr.StatusCode())
				}
				t.ExecuteTemplate(w, "base", statusErr.Error())
				return
			}
			http.Error(w, err.Error(), statusErr.StatusCode())
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
