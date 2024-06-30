package domain

type GETPriceDolar struct {
	Moneda    string  `json:"moneda"`
	Casa      string  `json:"casa"`
	Nombre    string  `json:"nombre"`
	Compra    float64 `json:"compra"`
	Venta     float64 `json:"venta"`
	JWT_PRICE string  `json:"jwtPrice"`
}
