package http

import (
	"fmt"

	"github.com/fxivan/set_up_server/microservice/internal/adapter/handler/request"
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
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
	var req request.CreateGiftCardRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, domain.ErrInternal)
		return
	}
	payload := getAuthPayload(ctx, authorizationPayloadKey)

	fmt.Print(payload)
	output, err := gc.svc.CreateGiftCardService(req, payload)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, output)
}