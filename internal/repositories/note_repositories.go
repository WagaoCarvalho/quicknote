package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/WagaoCarvalho/quicknote/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NoteRepository interface {
	ListNotes(ctx context.Context) ([]models.Note, error)
	GetNoteByID(ctx context.Context, id int) (*models.Note, error)
	CreateNote(ctx context.Context, title, content, color string) (*models.Note, error)
	UpdateNote(ctx context.Context, id int, title, content, color string) (*models.Note, error)
	DeleteNote(ctx context.Context, id int) error
}

func NewNoteRepository(dbPool *pgxpool.Pool) NoteRepository {
	return &noteRepository{
		db: dbPool,
	}
}

type noteRepository struct {
	db *pgxpool.Pool
}

func (nr *noteRepository) ListNotes(ctx context.Context) ([]models.Note, error) {
	var list []models.Note
	rows, err := nr.db.Query(ctx, "SELECT * FROM notes")
	if err != nil {
		return list, newRepositoryError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var note models.Note
		err = rows.Scan(
			&note.Id,
			&note.Title,
			&note.Content,
			&note.Color,
			&note.CreatedAt,
			&note.UpdatedAt,
		)
		if err != nil {
			return list, newRepositoryError(err)
		}
		list = append(list, note)
	}

	return list, nil
}

func (nr *noteRepository) GetNoteByID(ctx context.Context, id int) (*models.Note, error) {
	var note models.Note
	query := `SELECT id, title, content, color, created_at, updated_at 
			FROM notes WHERE id = $1`
	row := nr.db.QueryRow(ctx,
		query, id)
	if err := row.Scan(
		&note.Id,
		&note.Title,
		&note.Content,
		&note.Color,
		&note.CreatedAt,
		&note.UpdatedAt,
	); err != nil {
		return nil, newRepositoryError(err)
	}
	return &note, nil
}

func (nr *noteRepository) CreateNote(ctx context.Context, title, content, color string) (*models.Note, error) {
	var note models.Note

	query := `
		INSERT INTO notes (title, content, color)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	// Executa a query e preenche os campos do note
	err := nr.db.QueryRow(ctx, query, title, content, color).Scan(
		&note.Id,
		&note.CreatedAt,
	)
	if err != nil {
		return nil, newRepositoryError(err)
	}

	return &note, nil
}

func setValueIfNotEmpty(newValue, currentValue string) string {
	if newValue != "" {
		return newValue
	}
	return currentValue
}

func (nr *noteRepository) UpdateNote(ctx context.Context, id int, title, content, color string) (*models.Note, error) {
	var note models.Note
	note.Id = id
	note.Title = title
	note.Content = content
	note.Color = color

	note.UpdatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	query := `UPDATE notes set title = $1, content = $2, color = $3, updated_at = $4 
			  where id = $5`

	_, err := nr.db.Exec(ctx, query,
		note.Title, note.Content, note.Color,
		note.UpdatedAt, note.Id)
	if err != nil {
		return &note, newRepositoryError(err)
	}
	return &note, nil
}

func (nr *noteRepository) DeleteNote(ctx context.Context, id int) error {
	_, err := nr.db.Exec(ctx, "DELETE FROM notes WHERE id = $1", id)
	if err != nil {
		return newRepositoryError(err)
	}
	return nil
}
