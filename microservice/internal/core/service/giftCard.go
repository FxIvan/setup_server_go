package service

import (
	"github.com/fxivan/set_up_server/microservice/internal/adapter/config"
	"github.com/fxivan/set_up_server/microservice/internal/adapter/handler/request"
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"github.com/fxivan/set_up_server/microservice/internal/core/port"
)

type GiftCardService struct {
	repo port.UserService
	log  *config.TerminalLog
}

func NewGiftCardService(repo port.UserService, logTerminal *config.TerminalLog) *GiftCardService {
	return &GiftCardService{
		repo: repo,
		log:  logTerminal,
	}
}

func (gc *GiftCardService) CreateGiftCardService(body request.CreateGiftCardRequest, infoToken *domain.TokenPayload) (*domain.Coupon, error) {

	total := body.AmountCoupons * body.PriceCoupons

	coupon := &domain.Coupon{
		Owner:         infoToken.UserID,
		Title:         body.Title,
		Description:   body.Description,
		AmountCoupons: body.AmountCoupons,
		PriceCoupon:   body.PriceCoupons,
		Total:         total,
	}

	return coupon, nil
}
