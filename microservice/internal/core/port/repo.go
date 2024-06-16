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
	CreateNumberGiftCardStorage(amount int, collectionName string) ([]string, error)
	LinkingGiftCardUserStorage(collectionName string, coupons []string, infoPayment *domain.ResponseUalabisPOST, infoDomainCoupon *domain.Coupon) (string, error)
	//Payment
	SearchInfoPaymentStorage(collectionName string, idReference string) (*mongodb_model.CouponModel, error)
	UpdateStatusUala(collectionName string, idReference string, status string) error
}
