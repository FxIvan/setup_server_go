package dbRepository

import (
	"fmt"

	"github.com/fxivan/set_up_server/microservice/configuration"

	"github.com/fxivan/set_up_server/microservice/internal/adapter/config"
	mongodb "github.com/fxivan/set_up_server/microservice/internal/adapter/storage/mogodb"
	"github.com/fxivan/set_up_server/microservice/internal/adapter/storage/mysql"
	"github.com/fxivan/set_up_server/microservice/internal/adapter/storage/redis"
	"github.com/fxivan/set_up_server/microservice/internal/core/port"
)

func New(config *configuration.Configuration, logTerminal *config.TerminalLog) (port.RepoService, error) {
	var repo port.RepoService
	var err error

	switch config.Engine {
	case "mysql":
		repo, err = mysql.New(config, logTerminal)
	case "mongodb":
		repo, err = mongodb.New(config, logTerminal)
	case "redis":
		repo, err = redis.New(config, logTerminal)
	default:
		err = fmt.Errorf("invalid engine %s", config.Engine)
	}

	return repo, err
}
