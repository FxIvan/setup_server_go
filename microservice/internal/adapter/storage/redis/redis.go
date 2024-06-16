package redis

import (
	"context"
	"fmt"

	"github.com/fxivan/set_up_server/microservice/configuration"
	"github.com/fxivan/set_up_server/microservice/internal/adapter/config"
	mongodb_model "github.com/fxivan/set_up_server/microservice/internal/adapter/storage/mogodb/model"
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	connRedis *redis.Client
	log       *config.TerminalLog
}

func New(config *configuration.Configuration, logTerminal *config.TerminalLog) (*Redis, error) {
	ctx := context.Background()
	strConn := fmt.Sprintf("%s:%d", config.Host, config.Port)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     strConn,
		Password: "",
		DB:       0,
	})
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &Redis{
		connRedis: redisClient,
		log:       logTerminal,
	}, nil
}

func (m *Redis) CreateUserStorage(userModel *domain.User, collectionName string) (string, error) {
	return "", nil
}

func (m *Redis) GetUserEmailStorage(userEmail string, collectionName string) (*domain.User, error) {
	return nil, nil
}

func (m *Redis) ListUsersStorage(collectionName string) ([]domain.User, error) {
	return nil, nil
}

func (m *Redis) GetUserStorage(idUser string, collectionName string) (*domain.User, error) {
	return nil, nil
}

func (m *Redis) CreateNumberGiftCardStorage(amount int, collectionName string) ([]string, error) {
	return nil, nil
}

func (m *Redis) LinkingGiftCardUserStorage(collectionName string, coupons []string, infoPayment *domain.ResponseUalabisPOST, infoDomainCoupon *domain.Coupon) (string, error) {
	return "", nil
}

func (m *Redis) SearchInfoPaymentStorage(collectionName string, idReference string) (*mongodb_model.CouponModel, error) {
	return nil, nil
}
