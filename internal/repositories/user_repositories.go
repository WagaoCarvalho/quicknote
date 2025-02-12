package repositories

import (
	"context"

	"github.com/WagaoCarvalho/quicknote/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	ListNotes(ctx context.Context) ([]models.User, error)
	GetNoteByID(ctx context.Context, id int) (*models.Note, error)
	CreateNote(ctx context.Context, title, content, color string) (*models.Note, error)
	UpdateNote(ctx context.Context, id int, title, content, color string) (*models.Note, error)
	DeleteNote(ctx context.Context, id int) error
}

func NewUserRepository(dbPool *pgxpool.Pool) NoteRepository {
	return &noteRepository{
		db: dbPool,
	}
}

type userRepository struct {
	db *pgxpool.Pool
}
