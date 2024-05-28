package service

import (
	"fmt"

	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"github.com/fxivan/set_up_server/microservice/internal/core/port"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	db port.UserService
}

func NewUserService(db port.UserService) *UserService {
	return &UserService{db: db}
}

func (ur *UserService) FindUser(ctx *gin.Context, id string) (*domain.User, error) {
	salida, err := ur.db.Find(id)
	if err != nil {
		return nil, err
	}
	fmt.Print("salida", salida)
	userObj := &domain.User{
		ID:        salida.ID,
		Name:      salida.Name,
		Email:     salida.Email,
		Password:  salida.Password,
		CreatedAt: salida.CreatedAt,
		UpdatedAt: salida.UpdatedAt,
	}

	return userObj, nil
}
