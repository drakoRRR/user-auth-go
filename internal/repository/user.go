package repository

import (
	"context"
	"database/sql"
	"github.com/drakoRRR/user-auth-go/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (name, email, password, country) 
              VALUES ($1, $2, $3, $4) 
              RETURNING id`

	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Password, user.Country).Scan(&user.ID)
	return err
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, name, email, password, country, created_at, updated_at 
	          FROM users WHERE email = $1`
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password,
		&user.Country, &user.CreatedAt, &user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}
