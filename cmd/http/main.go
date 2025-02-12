package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/WagaoCarvalho/quicknote/internal/handlers"
	"github.com/WagaoCarvalho/quicknote/internal/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config := loadConfig()

	slog.SetDefault(newLogger(os.Stderr, config.GetLevelLog()))

	slog.Info(fmt.Sprintf("TESTE_LOAD_ENV: %s", config.TestLoadEnv))

	slog.Info(fmt.Sprintf("Server in 'http://127.0.0.1:%s'", config.ServerPort))

	dbUrl := config.DbConnUrl
	dbpool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("Connection success")
	defer dbpool.Close()

	mux := http.NewServeMux()

	staticHandler := http.FileServer(http.Dir("views/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", staticHandler))

	noteRepository := repositories.NewNoteRepository(dbpool)
	//userRepository := repositories.NewUserRepository(dbpool)

	noteHandler := handlers.NewNoteHandler(noteRepository)
	userHandler := handlers.NewUserHandler(nil)

	mux.HandleFunc("/", noteHandler.NotesList)

	mux.Handle("GET /note/{id}", handlers.HandlerWithError(noteHandler.NoteView))
	mux.Handle("GET /note/new", handlers.HandlerWithError(noteHandler.NoteNew))
	mux.Handle("POST /note", handlers.HandlerWithError(noteHandler.NoteSave))
	mux.Handle("DELETE /note/{id}", handlers.HandlerWithError(noteHandler.NoteDelete))
	mux.Handle("GET /note/{id}/edit", handlers.HandlerWithError(noteHandler.NoteEdit))

	mux.Handle("GET /user/signup", handlers.HandlerWithError(userHandler.SignupForm))
	mux.Handle("POST /user/signup", handlers.HandlerWithError(userHandler.Signup))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), mux); err != nil {
		panic(err)
	}
}
