package repository

import (
	"context"
	"log"
	"testing"

	"github.com/drakoRRR/user-auth-go/internal/models"
	"github.com/drakoRRR/user-auth-go/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	repo := repository.NewSQLUserRepository(testDB)

	ctx := context.Background()
	user := &models.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "hashed_password",
		Country:  "Testland",
	}

	err := repo.CreateUser(ctx, user)
	assert.NoError(t, err)

	retrievedUser, err := repo.GetUserByEmail(ctx, "test@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, retrievedUser)
	assert.Equal(t, "Test User", retrievedUser.Name)
	assert.Equal(t, "test@example.com", retrievedUser.Email)
}

func TestCreateUserWithEmptyFields(t *testing.T) {
	repo := repository.NewSQLUserRepository(testDB)

	ctx := context.Background()
	user := &models.User{
		Name:     "",
		Email:    "",
		Password: "hashed_password",
		Country:  "Testland",
	}

	err := repo.CreateUser(ctx, user)
	log.Println(err)
	assert.Error(t, err)
}

func TestCreateUserWithDuplicateEmail(t *testing.T) {
	repo := repository.NewSQLUserRepository(testDB)

	ctx := context.Background()
	user1 := &models.User{
		Name:     "User One",
		Email:    "duplicate@example.com",
		Password: "hashed_password",
		Country:  "Country A",
	}
	err := repo.CreateUser(ctx, user1)
	assert.NoError(t, err)

	user2 := &models.User{
		Name:     "User Two",
		Email:    "duplicate@example.com",
		Password: "hashed_password",
		Country:  "Country B",
	}
	err = repo.CreateUser(ctx, user2)
	assert.Error(t, err)
}

func TestGetUserByEmailNotFound(t *testing.T) {
	repo := repository.NewSQLUserRepository(testDB)

	ctx := context.Background()
	retrievedUser, err := repo.GetUserByEmail(ctx, "nonexistent@example.com")
	assert.NoError(t, err)
	assert.Nil(t, retrievedUser)
}

func TestGetUserByEmail(t *testing.T) {
	repo := repository.NewSQLUserRepository(testDB)

	ctx := context.Background()
	user := &models.User{
		Name:     "Test User",
		Email:    "test2@example.com",
		Password: "hashed_password",
		Country:  "Testland",
	}
	err := repo.CreateUser(ctx, user)
	assert.NoError(t, err)

	retrievedUser, err := repo.GetUserByEmail(ctx, "test@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, retrievedUser)
	assert.Equal(t, "Test User", retrievedUser.Name)
	assert.Equal(t, "test@example.com", retrievedUser.Email)
}
