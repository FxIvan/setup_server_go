package port

import (
	"context"

	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
)

type TokenService interface {
	CreateToken(userModel *domain.User) (string, error)
}

type AuthService interface {
	Login(ctx context.Context, email string, password string) (string, error)
}
