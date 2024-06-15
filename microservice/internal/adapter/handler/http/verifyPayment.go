package http

import (
	"github.com/fxivan/set_up_server/microservice/internal/core/service"
	"github.com/gin-gonic/gin"
)

type VerifyPaymentHandler struct {
	svc *service.VerifyPaymentService
}

func NewVerifyPaymentHandler(svc *service.VerifyPaymentService) *VerifyPaymentHandler {
	return &VerifyPaymentHandler{
		svc: svc,
	}
}

func (vp *VerifyPaymentHandler) VerifyPaymentHTTP(ctx *gin.Context) {
	uuidParam := ctx.Query("uuid")
	statusPayment := ctx.Query("status")
	err := vp.svc.UalaVerifyPaymentService(uuidParam, statusPayment)
	if err != nil {
		handleError(ctx, err)
		return
	}
	handleSuccess(ctx, "Payment Verify Success")
}
