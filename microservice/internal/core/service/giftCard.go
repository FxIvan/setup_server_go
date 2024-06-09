package service

import (
	"github.com/fxivan/set_up_server/microservice/internal/adapter/config"
	"github.com/fxivan/set_up_server/microservice/internal/adapter/handler/request"
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"github.com/fxivan/set_up_server/microservice/internal/core/port"
	"github.com/fxivan/set_up_server/microservice/internal/core/util"
)

type GiftCardService struct {
	repo port.UserService
	log  *config.TerminalLog
}

type BodyPaymentMicroservice struct {
	Amount         int    `json:"amount"`
	Description    string `json:"description"`
	SuccesResponse string `json:"succesResponse"`
	FailedResponse string `json:"failedResponse"`
}

func NewGiftCardService(repo port.UserService, logTerminal *config.TerminalLog) *GiftCardService {
	return &GiftCardService{
		repo: repo,
		log:  logTerminal,
	}
}

func (gc *GiftCardService) CreateGiftCardService(body request.CreateGiftCardRequest, infoToken *domain.TokenPayload) (*util.ResponsePOST, error) { //(*domain.Coupon, error)

	total := body.AmountCoupons * body.PriceCoupons

	coupon := &domain.Coupon{
		Owner:         infoToken.UserID,
		Title:         body.Title,
		Description:   body.Description,
		AmountCoupons: body.AmountCoupons,
		PriceCoupon:   body.PriceCoupons,
		Total:         total,
	}

	bodyPost := &BodyPaymentMicroservice{
		Amount:         coupon.Total,
		Description:    coupon.Description,
		SuccesResponse: "https://www.google.com/search?q=failed",
		FailedResponse: "https://www.google.com/search?q=success",
	}

	data, err := util.POSTMicroservice("http://localhost:3000/api/create/payment", "application/json ", bodyPost)
	if err != nil {
		return nil, err
	}

	return data, nil
}