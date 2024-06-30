package mongodb_model

import (
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
	Code   primitive.ObjectID `bson:"_id"`
	IdLink string             `bson:"IdLink"`
}

type CouponModel struct {
	IDReferentProcess string           `bson:"idDReferentProcess"`
	Email             string           `bson:"email"`
	Owner             string           `bson:"idOwner"`
	Title             string           `bson:"title"`
	Description       string           `bson:"description"`
	AmountCoupons     int              `bson:"amountCoupons"`
	PriceCoupon       float64          `bson:"priceCoupon"`
	Total             float64          `bson:"total"`
	Codes             []CouponMetaData `bson:"codes"`
	InfoPayment       LinkPaymentInfo  `bson:"infopayment"`
	CreatedAt         string           `bson:"createdAt"`
	UpdatedAt         string           `bson:"updatedAt"`
}

type CodeCoupon struct {
	ID                primitive.ObjectID `bson:"_id"`
	IDReferentProcess string             `bson:"idDReferentProcess"`
	UserOwner         string             `bson:"userOwner"`
	Code              string             `bson:"code"`
	BeneficiaryUser   string             `bson:"beneficiaryUser"`
	IsUsed            bool               `bson:"isUsed"`
	Price             float64            `bson:"price"`
	CVU               string             `bson:"cvu"`
	Alias             string             `bson:"alias"`
	Wallet            string             `bson:"wallet"`
	Red               string             `bson:"red"`
	CreatedAt         string             `bson:"createdAt"`
	UpdatedAt         string             `bson:"updatedAt"`
}
