package service

import (
	"context"

	"github.com/fxivan/set_up_server/microservice/internal/adapter/config"
	"github.com/fxivan/set_up_server/microservice/internal/core/port"
)

type AuthService struct {
	repo port.UserService
	ts   port.TokenService
	log  *config.TerminalLog
}

func NewAuthService(repo port.UserService, ts port.TokenService, logTerminal *config.TerminalLog) *AuthService {
	return &AuthService{
		repo: repo,
		ts:   ts,
		log:  logTerminal,
	}
}

func (au *AuthService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := au.repo.GetUserEmailStorage(email, "users")
	if err != nil {
		au.log.ErrorLog.Println(err)
		return "", err
	}

	accesToken, err := au.ts.CreateToken(user)
	if err != nil {
		au.log.ErrorLog.Println(err)
		return "", err
	}

	_, err = au.ts.VerifyToken(accesToken)
	if err != nil {
		return "", err
	}

	return accesToken, nil
}
