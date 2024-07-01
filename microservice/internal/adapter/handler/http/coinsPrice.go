package http

import (
	"github.com/fxivan/set_up_server/microservice/internal/core/service"
	"github.com/gin-gonic/gin"
)

type PriceCoinHandler struct {
	svc *service.CoinsPriceService
}

func NewPriceCoinHandler(svc *service.CoinsPriceService) *PriceCoinHandler {
	return &PriceCoinHandler{
		svc,
	}
}

func (pc *PriceCoinHandler) GetPriceDolarHTTP(ctx *gin.Context) {
	output, err := pc.svc.GetPriceDolarService()

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, "Precio dolar", output)
}
