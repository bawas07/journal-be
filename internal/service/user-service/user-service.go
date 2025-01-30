package userservice

import (
	baseservice "mindscribe-be/internal/service/base-service"
)

type UserService struct {
	base *baseservice.BaseService
}

func NewUserService(base *baseservice.BaseService) *UserService {
	return &UserService{
		base: base,
	}
}

func (s *UserService) CreateUser(email string, username string, password string) error {
	// Hash password
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return errors.New("Cannot encrypt password")
	// }

	// Create user
	// user := models.User{
	// 	Email:     email,
	// 	Username:  username,
	// 	Password:  string(hashedPassword),
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }
	return nil
}
