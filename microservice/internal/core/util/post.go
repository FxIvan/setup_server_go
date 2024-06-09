package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ResponsePOST struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func POSTMicroservice(URL string, format string, bodyPost any) (*ResponsePOST, error) {

	var responsePOST ResponsePOST
	jsonData, err := json.Marshal(bodyPost)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(URL, format, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

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

	// Procesar la respuesta (si es necesario)
	response := &ResponsePOST{
		Status:  responsePOST.Status,
		Message: responsePOST.Message,
		Data:    responsePOST.Data,
	}

	return response, nil

}
