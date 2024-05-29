package redis

import (
	"context"
	"fmt"

	"github.com/fxivan/set_up_server/microservice/configuration"
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	connRedis *redis.Client
}

func New(config *configuration.Configuration) (*Redis, error) {
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
	}, nil
}

func (m *Redis) CreateUserStorage(userModel *domain.User, collectionName string) (string, error) {
	return "", nil
}
