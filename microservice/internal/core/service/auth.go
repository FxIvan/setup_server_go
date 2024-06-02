package service

import (
	"context"

	"github.com/fxivan/set_up_server/microservice/internal/core/port"
)

type AuthService struct {
	repo port.UserService
	ts   port.TokenService
}

func NewAuthService(repo port.UserService, ts port.TokenService) *AuthService {
	return &AuthService{
		repo: repo,
		ts:   ts,
	}
}

func (au *AuthService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := au.repo.GetUserEmailStorage(email, "users")
	if err != nil {
		return "", err
	}

	accesToken, err := au.ts.CreateToken(user)
	if err != nil {
		return "", err
	}

	return accesToken, nil
}
