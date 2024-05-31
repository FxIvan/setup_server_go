package mongodb_model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModelMongoDB struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string
	Email     string `bson:"email"`
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
