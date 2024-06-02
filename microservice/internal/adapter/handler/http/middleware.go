package http

import (
	"strings"

	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"github.com/fxivan/set_up_server/microservice/internal/core/port"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationType       = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(token port.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		isEmpty := len(authorizationHeader) == 0
		if isEmpty {
			err := domain.ErrEmptyAuthorizationHeader
			handleAbort(ctx, err)
			return
		}

		fields := strings.Fields(authorizationHeader)
		isValid := len(fields) == 2
		if !isValid {
			err := domain.ErrInvalidAuthorizationHeader
			handleAbort(ctx, err)
			return
		}

		currentAuthorizationHeader := strings.ToLower(fields[0])
		if currentAuthorizationHeader != authorizationType {
			err := domain.ErrInvalidAuthorizationType
			handleAbort(ctx, err)
			return
		}

		accesToken := fields[1]
		payload, err := token.VerifyToken(accesToken)
		if err != nil {
			handleAbort(ctx, err)
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

func adminAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload := getAuthPayload(ctx, authorizationPayloadKey)
		isAdmin := payload.Role == domain.Admin
		if !isAdmin {
			err := domain.ErrForbidden
			handleAbort(ctx, err)
			return
		}
		ctx.Next()
	}
}
