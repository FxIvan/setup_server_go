package service

import (
	"github.com/fxivan/set_up_server/microservice/internal/adapter/config"
	"github.com/fxivan/set_up_server/microservice/internal/core/port"
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

func (vp *VerifyPaymentService) UalaVerifyPaymentService(uuid string, statusPayment string) error {

	//Buscamos el id de referencia creado por nosotros para buscar informacion del pago

	//Agarramos el ID del pago y buscamos en el miscroservicio
	/*res, err := util.GETVerifyPaymentUala(url)
	if err != nil {
		vp.log.ErrorLog.Println(err)
		return err
	}

	vp.log.InfoLog.Println("Response Verify Payment --->", res)*/

	return nil
}
