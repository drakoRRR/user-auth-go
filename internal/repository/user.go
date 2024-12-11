package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/drakoRRR/user-auth-go/internal/models"
	"github.com/drakoRRR/user-auth-go/pkg/utils"
)

// UserRepository defines the interface for user repository
type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

// SQLUserRepository is a concrete implementation of UserRepository
type SQLUserRepository struct {
	db *sql.DB
}

func NewSQLUserRepository(db *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{db: db}
}

func (r *SQLUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	if err := utils.ValidateUserData(user); err != nil {
		return err
	}

	query := `INSERT INTO users (name, email, password, country) 
              VALUES ($1, $2, $3, $4) 
              RETURNING id`

	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Password, user.Country).Scan(&user.ID)
	return err
}

func (r *SQLUserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}
	if !utils.IsValidEmail(email) {
		return nil, errors.New("invalid email format")
	}

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
