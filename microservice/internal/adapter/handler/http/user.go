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

type listUsersRequest struct {
	Skip  uint64 `form:"skip" binding:"required,min=0" example:"0"`
	Limit uint64 `form:"limit" binding:"required,min=5" example:"5"`
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

	handleSuccess(ctx, "Usuario registrado", text)
}

func (uh *UserHandler) ListUserHTTP(ctx *gin.Context) {
	var req listUsersRequest
	listUser, err := uh.svc.GetListUserService(ctx, req.Skip, req.Limit)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, "Usuarios", listUser)
}

func (uh *UserHandler) GetUserHTTP(ctx *gin.Context) {
	idUser := ctx.Param("id")

	user, err := uh.svc.GetUserService(ctx, idUser)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, "User", user)
}
