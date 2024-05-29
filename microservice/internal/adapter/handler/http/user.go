package http

/************************
Package: http
File: user.go
Arquitectura: Hexagonal
Description: Este paquete maneja las solicitudes HTTP relacionadas con los usuarios.
Flujo de datos: HTTP Request -> NeUserHandler(req,res) -> UserService(Controlador) -> UserRepository(DB)
*************************/

import (
	"time"

	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"github.com/fxivan/set_up_server/microservice/internal/core/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *service.UserService
}

type getUserRequest struct {
	ID string `uri:"id" binding:"required"`
}

type RegisterUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc,
	}
}

// FindUserByID retrieves a user by its ID
func (uh *UserHandler) FindUserByID(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		handleError(ctx, err)
		return
	}

	user, err := uh.svc.FindUser(ctx, req.ID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newUserResponse(user)
	handleSuccess(ctx, rsp)
}

func (uh *UserHandler) RegisterUser(ctx *gin.Context) {
	var user RegisterUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		handleError(ctx, err)
		return
	}

	userModel := &domain.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      domain.UserRole(user.Role),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	text, err := uh.svc.CreateUser(userModel)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, text)
}
