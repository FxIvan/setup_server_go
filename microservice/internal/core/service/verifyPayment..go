package service

import (
	"fmt"

	"github.com/fxivan/set_up_server/microservice/internal/adapter/config"
	"github.com/fxivan/set_up_server/microservice/internal/core/port"
	"github.com/fxivan/set_up_server/microservice/internal/core/util"
)

type VerifyPaymentService struct {
	repo port.RepoService
	log  *config.TerminalLog
}

func NewVerifyPaymentService(db port.RepoService, logTerminal *config.TerminalLog) *VerifyPaymentService {
	return &VerifyPaymentService{
		repo: db,
		log:  logTerminal,
	}
}

func (vp *VerifyPaymentService) UalaVerifyPaymentService(uuid string, statusPayment string) (string, error) {

	//Buscamos el id de referencia creado por nosotros para buscar informacion del pago
	infoPaymentDB, err := vp.repo.SearchInfoPaymentStorage("couponsalluser", uuid)

	if err != nil {
		vp.log.ErrorLog.Println(err)
		return "", err
	}

	//Agarramos el ID del pago y buscamos en el miscroservicio
	url := fmt.Sprintf("http://localhost:3000/api/verify/uala/%s", infoPaymentDB.InfoPayment.UUID)

	res, err := util.GETVerifyPaymentUala(url)
	if err != nil {
		vp.log.ErrorLog.Println(err)
		return "", err
	}

	//Aqui se debe actualizar es estado del PAGO en la base de datos
	err = vp.repo.UpdateStatusUala("couponsalluser", uuid, res.Data.Status)
	if err != nil {
		vp.log.ErrorLog.Println(err)
		return "", err
	}
	return res.Data.Status, nil
}
