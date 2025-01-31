package service

import (
	"context"
	"errors"
	"mindscribe-be/internal/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	base *BaseService
}

func newUserService(base *BaseService) *UserService {
	return &UserService{
		base: base,
	}
}

func (s *UserService) Create(ctx context.Context, email string, username string, password string) (error, *models.User) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Cannot encrypt password"), nil
	}

	// Create user
	user := models.User{
		Email:     email,
		Username:  username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.base.Repo.UserRepo.StoreUser(ctx, s.base.DB, &user)
	return err, &user
}
