package port

import "github.com/fxivan/set_up_server/microservice/internal/core/domain"

type UserService interface {
	Find(id string) (*domain.User, error)
	Save(data string) (*domain.User, error)
	FindKey(key string) (*domain.User, error)
}
