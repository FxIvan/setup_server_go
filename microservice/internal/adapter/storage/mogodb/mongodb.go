package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/fxivan/set_up_server/microservice/configuration"
	mongodbModel "github.com/fxivan/set_up_server/microservice/internal/adapter/storage/mogodb/model"
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

type User struct {
	ID        string    `json:"ID"`
	Name      string    `json:"Name"`
	Email     string    `json:"Email"`
	Password  string    `json:"Password"`
	Role      string    `json:"Role"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func New(config *configuration.Configuration) (*MongoDB, error) {

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=admin", config.User, config.Password, config.Host, config.Port, config.DBName)
	clientOptions := options.Client().ApplyURI(connectionString).SetAuth(options.Credential{
		Username: config.User,
		Password: config.Password,
	})
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("could not connect to MongoDB: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("could not ping MongoDB: %w", err)
	}

	db := client.Database(config.DBName)

	return &MongoDB{
		Client:   client,
		Database: db,
	}, nil
}

func (m *MongoDB) CreateUserStorage(userModel *domain.User, collectionName string) (string, error) {
	collection := m.Database.Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), userModel)
	if err != nil {
		return "", err
	}

	return "User Created", nil
}

func (m *MongoDB) GetUserEmailStorage(userEmail string, collectionName string) (*domain.User, error) {
	collection := m.Database.Collection(collectionName)
	var result mongodbModel.UserModelMongoDB
	filter := bson.M{"email": userEmail}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	userModelDomain := &domain.User{
		ID:        result.ID.Hex(),
		Name:      result.Name,
		Email:     result.Email,
		Password:  result.Password,
		Role:      domain.UserRole(result.Role),
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	return userModelDomain, nil
}
