package util

import (
	"github.com/google/uuid"
)

func GenerateUUIDUnique() string {
	newUUID := uuid.New()
	return newUUID.String()
}
