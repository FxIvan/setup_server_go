package domain

import (
	"github.com/google/uuid"
)

// UserRole is an entity that represents the role of the user

// TokenPayload is an entity that represents the payload of the token
type TokenPayload struct {
	ID     uuid.UUID
	UserID string
	Role   UserRole
}
