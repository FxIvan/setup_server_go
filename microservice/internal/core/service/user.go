package service

import (
	"time"

	"github.com/fxivan/set_up_server/microservice/internal/adapter/handler/request"
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"github.com/fxivan/set_up_server/microservice/internal/core/port"
)

type UserService struct {
	db port.UserService
}

func NewUserService(db port.UserService) *UserService {
	return &UserService{db: db}
}

func (m *UserService) CreateUserService(userModel *request.RegisterUserRequest) (string, error) {

	modelUser := &domain.User{
		Name:      userModel.Name,
		Email:     userModel.Email,
		Password:  userModel.Password,
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

func (m *UserService) LoginUserService(userModel *request.LoginUserRequest) (*domain.User, error) {
	modelUser := &domain.User{
		Email:    userModel.Email,
		Password: userModel.Password,
	}

	infoUser, err := m.db.GetUserEmailStorage(modelUser.Email, "users")
	if err != nil {
		return nil, err
	}

	return infoUser, nil
}
