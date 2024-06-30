package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
)

func GetPriceDolar(url string) (*domain.GETPriceDolar, error) {
	var responseDolar domain.GETPriceDolar
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer el cuerpo de la respuesta: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la respuesta del servidor: %s - %s", resp.Status, bodyResp)
	}

	if err := json.Unmarshal(bodyResp, &responseDolar); err != nil {
		return nil, fmt.Errorf("Error parsing request body: %s", err)
	}

	return &responseDolar, nil
}
