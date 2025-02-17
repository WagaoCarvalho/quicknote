package repositories

import (
	"context"
	"errors"

	"github.com/WagaoCarvalho/quicknote/internal/models"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrDuplicateEmail = newRepositoryError(errors.New("o e-mail já está em uso"))

type UserRepository interface {
	CreateUser(ctx context.Context, email, password string) (*models.User, error)
}

func NewUserRepository(dbPool *pgxpool.Pool) UserRepository {
	return &userRepository{
		db: dbPool,
	}
}

type userRepository struct {
	db *pgxpool.Pool
}

// CreateUser implements UserRepository.
func (ur *userRepository) CreateUser(ctx context.Context, email string, password string) (*models.User, error) {
	var user models.User

	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id, created_at
	`

	// Executa a query e preenche os campos do user
	err := ur.db.QueryRow(ctx, query, email, password).Scan(
		&user.Id,
		&user.CreatedAt,
	)
	if err != nil {
		// Verifica se o erro é uma violação de chave única (email já cadastrado)
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" { // Código 23505 = Unique Violation
			return nil, ErrDuplicateEmail
		}

		// Caso seja outro erro, retorna um erro genérico de repositório
		return nil, newRepositoryError(err)
	}

	return &user, nil
}
