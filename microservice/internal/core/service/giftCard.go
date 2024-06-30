package service

import (
	"fmt"

	"github.com/fxivan/set_up_server/microservice/internal/adapter/config"
	"github.com/fxivan/set_up_server/microservice/internal/adapter/handler/request"
	"github.com/fxivan/set_up_server/microservice/internal/adapter/handler/response"
	mongodb_model "github.com/fxivan/set_up_server/microservice/internal/adapter/storage/mogodb/model"
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"github.com/fxivan/set_up_server/microservice/internal/core/port"
	"github.com/fxivan/set_up_server/microservice/internal/core/util"
)

type GiftCardService struct {
	repo            port.RepoService
	log             *config.TerminalLog
	config          *config.URLMicroservice
	configContainer *config.Container
}

func NewGiftCardService(configLog *config.URLMicroservice, repo port.RepoService, logTerminal *config.TerminalLog, configContainer *config.Container) *GiftCardService {
	return &GiftCardService{
		repo:            repo,
		log:             logTerminal,
		config:          configLog,
		configContainer: configContainer,
	}
}

func (gc *GiftCardService) CreateGiftCardAuthService(body request.CreateGiftCardRequest, infoToken *domain.TokenPayload) (*response.ResCreatedGiftCard, error) {

	total := float64(body.AmountCoupons) * body.PriceCoupons

	coupon := &domain.Coupon{
		IDReference:   util.GenerateUUIDUnique(),
		Owner:         infoToken.UserID,
		Title:         body.Title,
		Description:   body.Description,
		AmountCoupons: body.AmountCoupons - 1,
		PriceCoupon:   body.PriceCoupons,
		Total:         total,
	}

	bodyPost := &request.RequestPaymentMicroservice{
		Amount:         coupon.Total,
		Description:    coupon.Description,
		SuccesResponse: fmt.Sprintf("https://api.tech/v1/verify/payment/uala?uuid=%s&status=success", coupon.IDReference),
		FailedResponse: fmt.Sprintf("https://api.tech/v1/verify/payment/uala?uuid=%s&status=failed", coupon.IDReference),
	}

	URLPost := fmt.Sprintf("%s/create/payment", gc.config.HostCreatePaymentNodeJS)

	data, err := util.POSTCreateGiftCardMicroservice(URLPost, "application/json ", bodyPost)
	if err != nil {
		gc.log.ErrorLog.Println(err)
		return nil, domain.ErrCreatedPaymentUala
	}

	allCode, err := gc.repo.CreateNumberGiftCardStorage(coupon.AmountCoupons, "coupons", infoToken, coupon)
	if err != nil {
		gc.log.ErrorLog.Println(err)
		return nil, domain.ErrCreatedNumberCoupons
	}

	bodyInfoPayment := &domain.ResponseUalabisPOST{
		IDReference: coupon.IDReference,
		IdTx:        data.Data.IdTx,
		Type:        data.Data.Type,
		UUID:        data.Data.UUID,
		OrderNumber: data.Data.OrderNumber,
		Amount:      data.Data.Amount,
		Status:      data.Data.Status,
		RefNumber:   data.Data.RefNumber,
		Links: domain.ResLinks{
			CheckoutLink: data.Data.Links.CheckoutLink,
			LinkSuccess:  data.Data.Links.LinkSuccess,
		},
	}

	dataLink, err := gc.repo.LinkingGiftCardUserStorage("couponsalluser", allCode, bodyInfoPayment, coupon)
	if err != nil {
		gc.log.ErrorLog.Println(err)
		return nil, domain.ErrLinkGiftCard
	}

	res := &response.ResCreatedGiftCard{
		Title:         dataLink.Title,
		Description:   dataLink.Description,
		AmountCoupons: dataLink.AmountCoupons,
		PriceCoupon:   dataLink.PriceCoupon,
		Total:         dataLink.Total,
	}

	return res, nil
}

func (gc *GiftCardService) GetGiftCardServicePublic(body request.CreateGiftCardRequest) (*response.ResCreatedGiftCard, error) {
	total := float64(body.AmountCoupons) * body.PriceCoupons

	priceDolarJWT, err := util.DecryptDolarPrice(body.JWTPriceDolar, gc.configContainer.JWT.JWT_SCRET)
	if err != nil {
		gc.log.ErrorLog.Println(err)
		return nil, domain.ErrDecryptPriceDolar
	}

	coupon := &domain.Coupon{
		IDReference:   util.GenerateUUIDUnique(),
		Owner:         "PUBLIC",
		Email:         body.Email,
		Title:         body.Title,
		Description:   body.Description,
		AmountCoupons: body.AmountCoupons - 1,
		PriceCoupon:   body.PriceCoupons * priceDolarJWT.Venta,
		Total:         total * priceDolarJWT.Venta,
	}

	bodyPost := &request.RequestPaymentMicroservice{
		Amount:         coupon.Total,
		Description:    coupon.Description,
		SuccesResponse: fmt.Sprintf("https://api.tech/v1/verify/payment/uala?uuid=%s&status=success", coupon.IDReference),
		FailedResponse: fmt.Sprintf("https://api.tech/v1/verify/payment/uala?uuid=%s&status=failed", coupon.IDReference),
	}

	URLPost := fmt.Sprintf("%s/create/payment", gc.config.HostCreatePaymentNodeJS)

	data, err := util.POSTCreateGiftCardMicroservice(URLPost, "application/json", bodyPost)
	if err != nil {
		gc.log.ErrorLog.Println(err)
		return nil, domain.ErrCreatedPaymentUala
	}

	infoToken := &domain.TokenPayload{
		UserID: "PUBLIC",
	}

	allCode, err := gc.repo.CreateNumberGiftCardStorage(coupon.AmountCoupons, "coupons", infoToken, coupon)
	if err != nil {
		gc.log.ErrorLog.Println(err)
		return nil, domain.ErrCreatedNumberCoupons
	}

	bodyInfoPayment := &domain.ResponseUalabisPOST{
		IDReference: coupon.IDReference,
		IdTx:        data.Data.IdTx,
		Type:        data.Data.Type,
		UUID:        data.Data.UUID,
		OrderNumber: data.Data.OrderNumber,
		Amount:      data.Data.Amount,
		Status:      data.Data.Status,
		RefNumber:   data.Data.RefNumber,
		Links: domain.ResLinks{
			CheckoutLink: data.Data.Links.CheckoutLink,
			LinkSuccess:  data.Data.Links.LinkSuccess,
		},
	}

	dataLink, err := gc.repo.LinkingGiftCardUserStorage("couponsalluser", allCode, bodyInfoPayment, coupon)
	if err != nil {
		gc.log.ErrorLog.Println(err)
		return nil, domain.ErrLinkGiftCard
	}

	res := &response.ResCreatedGiftCard{
		Title:         dataLink.Title,
		Description:   dataLink.Description,
		AmountCoupons: dataLink.AmountCoupons,
		PriceCoupon:   dataLink.PriceCoupon,
		Total:         dataLink.Total,
	}

	return res, nil
}

func (gc *GiftCardService) InsertCodeService(body request.InsertCodeRequest) (*mongodb_model.CodeCoupon, error) {

	couponModel, err := gc.repo.SearchCodeStorage("coupons", body.Code)
	if err != nil {
		return nil, domain.ErrSearchCode
	}

	userAndCoupons, err := gc.repo.SearchCouponsAllUser("couponsalluser", couponModel.IDReferentProcess)
	if err != nil {
		return nil, domain.ErrSearchCode
	}

	isAPPROVED := userAndCoupons.InfoPayment.Status
	if isAPPROVED != "APPROVED" {
		return nil, domain.ErrBottomlessCoupon
	}

	couponToUpdated := &mongodb_model.CodeCoupon{
		IsUsed: true,
		CVU:    body.CVU,
		Alias:  body.Alias,
		Wallet: body.Wallet,
		Red:    body.Red,
	}

	couponUpdated, err := gc.repo.UpdateCouponStorage("coupons", couponToUpdated, body.Code)
	if err != nil {
		return nil, domain.ErrSearchCode
	}

	return couponUpdated, nil
}
