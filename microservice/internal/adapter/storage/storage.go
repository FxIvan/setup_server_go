package dbRepository

import (
	"fmt"

	"github.com/fxivan/set_up_server/microservice/configuration"

	mongodb "github.com/fxivan/set_up_server/microservice/internal/adapter/storage/mogodb"
	"github.com/fxivan/set_up_server/microservice/internal/adapter/storage/mysql"
	"github.com/fxivan/set_up_server/microservice/internal/adapter/storage/redis"
	"github.com/fxivan/set_up_server/microservice/internal/core/port"
)

func New(config *configuration.Configuration) (port.UserService, error) {
	var repo port.UserService
	var err error

	switch config.Engine {
	case "mysql":
		repo, err = mysql.New(config)
	case "mongodb":
		repo, err = mongodb.New(config)
	case "redis":
		repo, err = redis.New(config)
	default:
		err = fmt.Errorf("invalid engine %s", config.Engine)
	}

	return repo, err
}
