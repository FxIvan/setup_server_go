package http

import (
	"fmt"

	"github.com/fxivan/set_up_server/microservice/internal/core/service"
	"github.com/gin-gonic/gin"
)

type GiftCardHandler struct {
	svc *service.GiftCardService
}

func NewGiftCardHandler(svc *service.GiftCardService) *GiftCardHandler {
	return &GiftCardHandler{
		svc,
	}
}

func (gc *GiftCardHandler) CreateGiftCardHTTP(ctx *gin.Context) {

	payload := getAuthPayload(ctx, authorizationPayloadKey)
	fmt.Print(payload)
}
