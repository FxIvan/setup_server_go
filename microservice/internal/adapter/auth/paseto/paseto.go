package paseto

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/fxivan/set_up_server/microservice/internal/adapter/config"
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"github.com/fxivan/set_up_server/microservice/internal/core/port"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type PasetoToken struct {
	token      *paseto.JSONToken
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
	duration   time.Duration
}

func New(config *config.Token) (port.TokenService, error) {
	durationStr := config.Duration
	privateKey := config.PrivateKey
	publicKey := config.PublicKey

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return nil, domain.ErrTokenDuration
	}

	issuedAt := time.Now()
	expireAt := issuedAt.Add(duration)

	token := paseto.JSONToken{
		Expiration: expireAt,
		IssuedAt:   issuedAt,
		NotBefore:  expireAt,
	}

	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}

	publicKeyBytes, err := hex.DecodeString(publicKey)
	if err != nil {
		return nil, err
	}

	parsePrivateKey := ed25519.PrivateKey(privateKeyBytes)
	parsePublicKey := ed25519.PublicKey(publicKeyBytes)

	return &PasetoToken{
		token:      &token,
		privateKey: parsePrivateKey,
		publicKey:  parsePublicKey,
		duration:   duration,
	}, nil
}

func (pt *PasetoToken) CreateToken(userModel *domain.User) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return "", domain.ErrTokenCreation
	}

	payload := &domain.TokenPayload{
		ID:     id,
		UserID: userModel.ID,
		Role:   userModel.Role,
	}

	pt.token.Set("id", payload.ID.String())
	pt.token.Set("userID", payload.UserID)
	pt.token.Set("role", string(payload.Role))
	footer := "some footer"
	v2 := paseto.NewV2()
	fmt.Println(pt.privateKey)

	//Sign data
	token, err := v2.Sign(pt.privateKey, *pt.token, paseto.ParseFooter(footer, nil))
	if err != nil {
		fmt.Println(err)
		return "", domain.ErrTokenCreation
	}

	return token, nil
}
