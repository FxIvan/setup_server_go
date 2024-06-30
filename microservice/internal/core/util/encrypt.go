package util

import (
	"fmt"
	"time"

	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"github.com/golang-jwt/jwt/v5"
)

type EncryptDolarPriceModel struct {
	Moneda string  `json:"moneda"`
	Casa   string  `json:"casa"`
	Nombre string  `json:"nombre"`
	Compra float64 `json:"compra"`
	Venta  float64 `json:"venta"`
	jwt.RegisteredClaims
}

func EncryptDolarPrice(body *domain.GETPriceDolar, secret string) (*domain.GETPriceDolar, error) {

	fmt.Println(body)
	claims := EncryptDolarPriceModel{
		body.Moneda,
		body.Casa,
		body.Nombre,
		body.Compra,
		body.Venta,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
			Issuer:    "PriceDolar",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	body.JWT_PRICE = tokenStr

	return body, nil
}

func DecryptDolarPrice(tokenStr string, secret string) (*domain.GETPriceDolar, error) {
	claims := EncryptDolarPriceModel{}

	_, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	body := domain.GETPriceDolar{
		Moneda:    claims.Moneda,
		Casa:      claims.Casa,
		Nombre:    claims.Nombre,
		Compra:    claims.Compra,
		Venta:     claims.Venta,
		JWT_PRICE: tokenStr,
	}

	return &body, nil
}
