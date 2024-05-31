package port

import (
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
)

type UserService interface {
	CreateUserStorage(userModel *domain.User, collectionName string) (string, error)
	GetUserEmailStorage(userEmail string, collectionName string) (*domain.User, error)
}
