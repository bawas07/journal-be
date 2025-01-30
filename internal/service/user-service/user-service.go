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
