package service

import (
	"github.com/fxivan/set_up_server/microservice/internal/adapter/config"
	"github.com/fxivan/set_up_server/microservice/internal/core/port"
)

type GiftCardService struct {
	repo port.UserService
	log  *config.TerminalLog
}

func NewGiftCardService(repo port.UserService, logTerminal *config.TerminalLog) *GiftCardService {
	return &GiftCardService{
		repo: repo,
		log:  logTerminal,
	}
}
