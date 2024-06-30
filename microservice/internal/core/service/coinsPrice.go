package service

import (
	"github.com/fxivan/set_up_server/microservice/internal/adapter/config"
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"github.com/fxivan/set_up_server/microservice/internal/core/port"
	"github.com/fxivan/set_up_server/microservice/internal/core/util"
)

type CoinsPriceService struct {
	repo   port.RepoService
	log    *config.TerminalLog
	config *config.Container
}

func NewCoinsPriceService(configEnv *config.Container, repo port.RepoService, logTerminal *config.TerminalLog) *CoinsPriceService {
	return &CoinsPriceService{
		repo:   repo,
		log:    logTerminal,
		config: configEnv,
	}
}

func (cp *CoinsPriceService) GetPriceDolarService() (*domain.GETPriceDolar, error) {
	response, err := util.GetPriceDolar("https://dolarapi.com/v1/dolares/tarjeta")
	if err != nil {
		cp.log.ErrorLog.Println(err)
		return nil, domain.ErrDataNotFound
	}

	encrypt, err := util.EncryptDolarPrice(response, cp.config.JWT.JWT_SCRET)

	if err != nil {
		cp.log.ErrorLog.Println(err)
		return nil, domain.ErrEncryptData
	}

	return encrypt, nil

}
