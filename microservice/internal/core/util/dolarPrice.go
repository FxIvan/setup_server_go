package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ReponseDolarPrice struct {
	Moneda string  `json:"moneda"`
	Casa   string  `json:"casa"`
	Nombre string  `json:"nombre"`
	Compra float64 `json:"compra"`
	Venta  float64 `json:"venta"`
}

func GetPriceDolar(url string) (*ReponseDolarPrice, error) {
	var responseDolar ReponseDolarPrice
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
