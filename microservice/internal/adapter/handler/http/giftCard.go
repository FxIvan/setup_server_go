package http

import (
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

func (gc *GiftCardHandler) CreateGiftCardAuthHTTP(ctx *gin.Context) {
	var req request.CreateGiftCardRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, domain.ErrInternal)
		return
	}
	payload := getAuthPayload(ctx, authorizationPayloadKey)

	output, err := gc.svc.CreateGiftCardAuthService(req, payload)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, output)
}

func (gc *GiftCardHandler) CreateGiftCardPublicHTTP(ctx *gin.Context) {
	var req request.CreateGiftCardRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, domain.ErrInternal)
		return
	}
	output, err := gc.svc.GetGiftCardServicePublic(req)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, output)
}

func (gc *GiftCardHandler) SearchCodeHTTP(ctx *gin.Context) {
	var req request.InsertCodeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, domain.ErrInternal)
		return
	}

	res, err := gc.svc.InsertCodeService(req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, res)
}
