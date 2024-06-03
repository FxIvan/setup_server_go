package service

import (
	"context"
	"time"

	"github.com/fxivan/set_up_server/microservice/internal/adapter/handler/request"
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"github.com/fxivan/set_up_server/microservice/internal/core/port"
	"github.com/fxivan/set_up_server/microservice/internal/core/util"
)

type UserService struct {
	db port.UserService
}

func NewUserService(db port.UserService) *UserService {
	return &UserService{db: db}
}

func (m *UserService) CreateUserService(userModel *request.RegisterUserRequest) (string, error) {

	hashedPassword, err := util.HashPassword(userModel.Password)
	if err != nil {
		return "", err
	}

	modelUser := &domain.User{
		Name:      userModel.Name,
		Email:     userModel.Email,
		Password:  hashedPassword,
		Role:      domain.UserRole(userModel.Role),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	salida, err := m.db.CreateUserStorage(modelUser, "users")
	if err != nil {
		return "", err
	}

	return salida, nil
}

func (m *UserService) GetListUserService(ctx context.Context, skip, limit uint64) ([]domain.User, error) {
	var users []domain.User

	users, err := m.db.ListUsersStorage("users")
	if err != nil {
		return nil, domain.ErrDataNotFound
	}

	return users, nil
}

func (m *UserService) GetUserService(ctx context.Context, id string) (*domain.User, error) {
	user, err := m.db.GetUserStorage(id, "users")
	if err != nil {
		return nil, domain.ErrDataNotFound
	}

	return user, nil
}
