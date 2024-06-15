package mongodb_model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BeneficiaryUserInfo struct {
	ID     primitive.ObjectID `bson:"userBeneficiaryID"`
	Name   string
	Code   int
	Mobile int
	Email  string
}

type LinkPaymentInfo struct {
	Link        string
	SuccessLink string
	FailedLink  string
	OrderNumber string
	Amount      string
	RefNumber   string
	Status      string
	Type        string
	IdTx        string
	UUID        string
}

type CouponMetaData struct {
	Code            string
	ExpireAt        time.Time
	BeneficiaryUser BeneficiaryUserInfo
	IsUsed          bool
	Price           int
	CVU             string
	Alias           string
	Wallet          string
	Red             string
}

type CouponModel struct {
	IDReferentProcess string `bson:"idDReferentProcess"`
	Owner             string `bson:"idOwner"`
	Title             string `bson:"title"`
	Description       string `bson:"description"`
	AmountCoupons     string `bson:"amountCoupons"`
	PriceCoupon       string `bson:"priceCoupon"`
	Total             string `bson:"total"`
	Codes             []CouponMetaData
	InfoPayment       LinkPaymentInfo
}

type CodeCoupon struct {
	Code string
}
