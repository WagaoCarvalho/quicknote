package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

func notes(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"views/templates/base.html",
		"views/templates/pages/home.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Erro ao tentar acessar a página: "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", nil)
}

func note(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // Obtendo o id da URL
	if id == "" {
		http.Error(w, "O id é obrigatório", http.StatusNotFound)
		return
	}

	files := []string{
		"views/templates/base.html",
		"views/templates/pages/note.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Erro ao tentar acessar a página: "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", id)
}

func noteNew(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/note_new.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Erro ao tentar acessar a página: "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", nil)
}

func noteCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(405)
		http.Error(w, "Método http não permitido", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("noteCreate"))
	fmt.Fprint(w, "Criando...")
}

func main() {
	config := loadConfig()
	slog.SetDefault(newLogger(os.Stderr, config.GetLevelLog()))

	slog.Info(fmt.Sprintf("TESTE_LOAD_ENV: %s", config.TestLoadEnv))

	slog.Info(fmt.Sprintf("Server in 'http://127.0.0.1:%s'", config.ServerPort))

	mux := http.NewServeMux()

	staticHandler := http.FileServer(http.Dir("views/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", staticHandler))

	mux.HandleFunc("/", notes)
	mux.HandleFunc("/note/{id}/view", note)
	mux.HandleFunc("/note/new", noteNew)
	mux.HandleFunc("/note/create", noteCreate)

	http.Handle("/", mux)
	fmt.Println("Server in http://127.0.0.1:5000")
	http.ListenAndServe(":5000", mux)
}
