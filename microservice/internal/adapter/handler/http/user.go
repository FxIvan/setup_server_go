package http

/************************
Package: http
File: user.go
Arquitectura: Hexagonal
Description: Este paquete maneja las solicitudes HTTP relacionadas con los usuarios.
Flujo de datos: HTTP Request -> NeUserHandler(req,res) -> UserService(Controlador) -> UserRepository(DB)
*************************/

import (
	"github.com/fxivan/set_up_server/microservice/internal/adapter/handler/request"
	"github.com/fxivan/set_up_server/microservice/internal/core/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *service.UserService
}

type getUserRequest struct {
	ID string `uri:"id" binding:"required"`
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc,
	}
}

// Registro de usuario HTTP
func (uh *UserHandler) RegisterUserHTTP(ctx *gin.Context) {
	var user request.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		handleError(ctx, err)
		return
	}
	text, err := uh.svc.CreateUserService(&user)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, text)
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
