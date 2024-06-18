package mongodb_model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BeneficiaryUserInfo struct {
	Name   string `bson:"name"`
	Code   string `bson:"code"`
	Mobile string `bson:"mobile"`
	Email  string `bson:"email"`
}

type LinkPaymentInfo struct {
	Link        string `bson:"link"`
	SuccessLink string `bson:"successlink"`
	FailedLink  string `bson:"failedlink"`
	OrderNumber string `bson:"ordernumber"`
	Amount      string `bson:"amount"`
	RefNumber   string `bson:"refnumber"`
	Status      string `bson:"status"`
	Type        string `bson:"type"`
	IdTx        string `bson:"idtx"`
	UUID        string `bson:"uuid"`
}

type CouponMetaData struct {
	Code            primitive.ObjectID  `bson:"_id"`
	ExpireAt        time.Time           `bson:"expireat"`
	BeneficiaryUser BeneficiaryUserInfo `bson:"beneficiaryuser"`
}

type CouponModel struct {
	IDReferentProcess string           `bson:"idDReferentProcess"`
	Owner             string           `bson:"idOwner"`
	Title             string           `bson:"title"`
	Description       string           `bson:"description"`
	AmountCoupons     string           `bson:"amountCoupons"`
	PriceCoupon       string           `bson:"priceCoupon"`
	Total             string           `bson:"total"`
	Codes             []CouponMetaData `bson:"codes"`
	InfoPayment       LinkPaymentInfo  `bson:"infopayment"`
}

type CodeCoupon struct {
	ID              primitive.ObjectID `bson:"_id"`
	UserOwner       string             `bson:"userOwner"`
	Code            string             `bson:"code"`
	BeneficiaryUser string             `bson:"beneficiaryUser"`
	IsUsed          bool               `bson:"isUsed"`
	Price           int                `bson:"price"`
	CVU             string             `bson:"cvu"`
	Alias           string             `bson:"alias"`
	Wallet          string             `bson:"wallet"`
	Red             string             `bson:"red"`
}
