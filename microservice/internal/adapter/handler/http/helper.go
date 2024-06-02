package http

import (
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func getAuthPayload(ctx *gin.Context, key string) *domain.TokenPayload {
	return ctx.MustGet(key).(*domain.TokenPayload)
}
