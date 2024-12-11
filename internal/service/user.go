package service

import (
	"context"
	"errors"
	"github.com/drakoRRR/user-auth-go/internal/auth"
	"github.com/drakoRRR/user-auth-go/internal/ipgeolocation"
	"github.com/drakoRRR/user-auth-go/internal/models"
	"github.com/drakoRRR/user-auth-go/internal/repository"
	"github.com/go-playground/validator/v10"
	"log"
	"time"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, user models.RegisterUserPayload, ipAddress string) (*models.UserResponse, error)
}

type UserService struct {
	repo      repository.UserRepository
	validator *validator.Validate
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo:      repo,
		validator: validator.New(),
	}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	user models.RegisterUserPayload,
	ipAddress string,
) (*models.UserResponse, error) {
	if err := s.validator.Struct(user); err != nil {
		return nil, err
	}

	existingUser, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	country, err := ipgeolocation.GetCountry(ipAddress)
	if err != nil {
		log.Printf("Error enriching country for IP %s: %v", ipAddress, err)
		country = "Unknown"
	}

	newUser := models.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  hashedPassword,
		Country:   country,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.repo.CreateUser(ctx, &newUser)
	if err != nil {
		return nil, err
	}

	userResponse := &models.UserResponse{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		Country:   newUser.Country,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}

	return userResponse, nil
}
