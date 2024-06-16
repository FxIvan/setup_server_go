package port

import (
	"context"

	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
)

type TokenService interface {
	CreateToken(userModel *domain.User) (string, error)
	VerifyToken(token string) (*domain.TokenPayload, error)
}

type AuthService interface {
	LoginService(ctx context.Context, email string, password string) (string, error)
}
