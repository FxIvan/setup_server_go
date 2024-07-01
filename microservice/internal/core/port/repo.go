package port

import (
	mongodb_model "github.com/fxivan/set_up_server/microservice/internal/adapter/storage/mogodb/model"
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
)

type RepoService interface {
	//User
	CreateUserStorage(userModel *domain.User, collectionName string) (string, error)
	GetUserEmailStorage(userEmail string, collectionName string) (*domain.User, error)
	ListUsersStorage(collectionName string) ([]domain.User, error)
	GetUserStorage(idUser string, collectionName string) (*domain.User, error)
	//GiftCard
	CreateNumberGiftCardStorage(amount int, collectionName string, infoToken *domain.TokenPayload, infoCoupon *domain.Coupon) ([]mongodb_model.CodeCoupon, error)
	LinkingGiftCardUserStorage(collectionName string, coupons []mongodb_model.CodeCoupon, infoPayment *domain.ResponseUalabisPOST, infoDomainCoupon *domain.Coupon) (*mongodb_model.CouponModel, error)
	SearchCodeStorage(collectionName string, codeName string) (*mongodb_model.CodeCoupon, error)
	UpdateCouponStorage(collectionName string, couponUpdated *mongodb_model.CodeCoupon, codeName string) (*mongodb_model.CodeCoupon, error)
	SearchCouponsAllUser(collectionName string, idReference string) (*mongodb_model.CouponModel, error)
	//Payment
	SearchInfoPaymentStorage(collectionName string, idReference string) (*mongodb_model.CouponModel, error)
	UpdateStatusUalaStorage(collectionName string, idReference string, status string) error
}
