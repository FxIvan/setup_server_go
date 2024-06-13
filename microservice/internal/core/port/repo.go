package port

import (
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
	LinkingGiftCardUserStorage(collectionName string, coupons []string, infoPayment any, infoDomainCoupon *domain.Coupon) (string, error)
}
