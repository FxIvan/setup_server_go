package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
)

type ResponsePOST struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    domain.ResponseUalabisPOST `json:"data"`
}

func POSTCreateGiftCardMicroservice(URL string, format string, bodyPost any) (*ResponsePOST, error) {
	var responsePOST ResponsePOST
	jsonData, err := json.Marshal(bodyPost)
	if err != nil {
		return nil, err
	}

	start := time.Now() // Start time

	resp, err := http.Post(URL, format, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	duration := time.Since(start) // Calculate duration
	fmt.Printf("POST request took %s\n", duration)

	// Leer el cuerpo de la respuesta
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer el cuerpo de la respuesta: %v", err)
	}

	// Verificar el código de estado de la respuesta
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la respuesta del servidor: %s - %s", resp.Status, bodyResp)
	}

	if err := json.Unmarshal(bodyResp, &responsePOST); err != nil {
		return nil, fmt.Errorf("Error parsing request body: %s", err)
	}

	responsePOSTUalabis := &domain.ResponseUalabisPOST{
		IdTx:        responsePOST.Data.IdTx,
		Type:        responsePOST.Data.Type,
		UUID:        responsePOST.Data.UUID,
		OrderNumber: responsePOST.Data.OrderNumber,
		Amount:      responsePOST.Data.Amount,
		Status:      responsePOST.Data.Status,
		RefNumber:   responsePOST.Data.RefNumber,
		Links: domain.ResLinks{
			CheckoutLink: responsePOST.Data.Links.CheckoutLink,
			LinkSuccess:  responsePOST.Data.Links.LinkSuccess,
			LinkFailed:   responsePOST.Data.Links.LinkFailed,
		},
	}

	responsePOST.Data = *responsePOSTUalabis

	return &responsePOST, nil

}

func GETVerifyPaymentUala(URL string) (*domain.ResponseUalabisPOSTVerify, error) {
	var resPOST domain.ResponseUalabisPOSTVerify

	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer el cuerpo de la respuesta: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la respuesta del servidor: %s - %s", resp.Status, bodyResp)
	}

	if err := json.Unmarshal(bodyResp, &resPOST); err != nil {
		return nil, fmt.Errorf("Error parsing request body: %s", err)
	}

	return &resPOST, nil
}
