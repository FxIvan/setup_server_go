package http

import (
	"github.com/fxivan/set_up_server/microservice/internal/adapter/handler/request"
	"github.com/fxivan/set_up_server/microservice/internal/core/port"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc port.AuthService
}

func NewAuthHandler(svc port.AuthService) *AuthHandler {
	return &AuthHandler{
		svc,
	}
}

func (au *AuthHandler) LoginUserHTTP(ctx *gin.Context) {
	var user request.LoginUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		handleError(ctx, err)
		return
	}

	token, err := au.svc.LoginService(ctx, user.Email, user.Password)
	if err != nil {
		handleError(ctx, err)
		return

	}
	handleSuccess(ctx, "Login inciado", token)
}
