package handlers

import (
	"html/template"
	"net/http"
)

func render(w http.ResponseWriter, statusCode int, page string, data interface{}) {
	// Lista de arquivos para o template
	filesPages := []string{
		"views/templates/base.html",
		"views/templates/pages/" + page,
	}

	// Parse dos arquivos de template
	t, err := template.ParseFiles(filesPages...)
	if err != nil {
		http.Error(w, "Erro ao carregar o template", http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho de status
	w.WriteHeader(statusCode)

	// Executa o template com os dados fornecidos
	err = t.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Erro ao renderizar a página", http.StatusInternalServerError)
		return
	}
}
