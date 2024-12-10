package service_test

import (
	"context"
	"testing"

	"github.com/drakoRRR/user-auth-go/internal/models"
	"github.com/drakoRRR/user-auth-go/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) CreateUser(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*models.User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepo)
	userService := service.NewUserService(mockRepo)

	user := models.RegisterUserPayload{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
	}

	// Mock GetUserByEmail to return nil (no existing user)
	mockRepo.On("GetUserByEmail", mock.Anything, user.Email).Return((*models.User)(nil), nil)
	// Mock CreateUser to return nil (successful creation)
	mockRepo.On("CreateUser", mock.Anything, mock.Anything).Return(nil)

	result, err := userService.CreateUser(context.Background(), user, "127.0.0.1")

	assert.NoError(t, err)
	assert.Equal(t, user.Email, result.Email)
	assert.Equal(t, user.Name, result.Name)

	mockRepo.AssertExpectations(t)
}

func TestCreateUserWithEmptyName(t *testing.T) {
	mockRepo := new(MockUserRepo)
	userService := service.NewUserService(mockRepo)

	user := models.RegisterUserPayload{
		Name:     "",
		Email:    "test@example.com",
		Password: "password123",
	}

	// Try to create user with empty name, expect validation error
	result, err := userService.CreateUser(context.Background(), user, "127.0.0.1")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Field validation for 'Name' failed on the 'required' tag")

	mockRepo.AssertExpectations(t)
}

func TestCreateUserWithDuplicateEmail(t *testing.T) {
	mockRepo := new(MockUserRepo)
	userService := service.NewUserService(mockRepo)

	user := models.RegisterUserPayload{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
	}

	// Mock GetUserByEmail to return a user (duplicate email)
	existingUser := &models.User{
		ID:    "123",
		Name:  "Test User",
		Email: "test@example.com",
	}
	mockRepo.On("GetUserByEmail", mock.Anything, user.Email).Return(existingUser, nil)

	result, err := userService.CreateUser(context.Background(), user, "127.0.0.1")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "user with this email already exists", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestCreateUserWithInvalidEmail(t *testing.T) {
	mockRepo := new(MockUserRepo)
	userService := service.NewUserService(mockRepo)

	user := models.RegisterUserPayload{
		Name:     "Test User",
		Email:    "invalid-email",
		Password: "password123",
	}

	// Try to create user with invalid email, expect validation error
	result, err := userService.CreateUser(context.Background(), user, "127.0.0.1")

	assert.Error(t, err)
	assert.Nil(t, result)

	mockRepo.AssertExpectations(t)
}
