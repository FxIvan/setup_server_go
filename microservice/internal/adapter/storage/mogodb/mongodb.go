package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/fxivan/set_up_server/microservice/configuration"
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
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
	collection := db.Collection("microservicio")

	return &MongoDB{
		Client:     client,
		Database:   db,
		Collection: collection,
	}, nil
}

func (m *MongoDB) Find(id_user string) (*domain.User, error) {

	jsonData := `
	{
		"a1DQAsd2q1dd5d1wds1d2w": {
			"ID": "a1DQAsd2q1dd5d1wds1d2w",
			"Name": "John Doe",
			"Email": "jhon@gmail.com",
			"Password": "password",
			"Role": "admin",
			"CreatedAt": "2021-09-01T00:00:00Z",
			"UpdatedAt": "2021-09-01T00:00:00Z"
		},
		"91DQAsd2q1dd5d1wds1dty": {
			"ID": "91DQAsd2q1dd5d1wds1dty",
			"Name": "John Doe",
			"Email": "doe@gmail.com",
			"Password": "password",
			"Role": "user",
			"CreatedAt": "2021-09-01T00:00:00Z",
			"UpdatedAt": "2021-09-01T00:00:00Z"
		},
		"81DQAsd2q1dd5d1wds9ePo": {
			"ID": "81DQAsd2q1dd5d1wds9ePo",
			"Name": "John Doe",
			"Email": "gmail@gmail.com",
			"Password": "password",
			"Role": "user",
			"CreatedAt": "2021-09-01T00:00:00Z",
			"UpdatedAt": "2021-09-01T00:00:00Z"
		}
	}
	
	`
	// Map to hold the decoded JSON
	users := make(map[string]User)

	// Unmarshal the JSON data
	err := json.Unmarshal([]byte(jsonData), &users)
	if err != nil {
		fmt.Printf("failed to unmarshal JSON data: %v\n", err)
		return nil, err
	}

	// Print the decoded users
	for _, user := range users {
		//fmt.Printf("ID: %s, User: %+v\n", id, user)
		if id_user == user.ID {
			response := &domain.User{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				Password:  user.Password,
				Role:      domain.UserRole(user.Role),
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			}
			return response, nil
		}
	}

	return nil, nil
}

func (r *MongoDB) FindKey(key string) (*domain.User, error) {
	return nil, nil
}

func (m *MongoDB) Save(data string) (*domain.User, error) {
	_, err := m.Collection.InsertOne(context.Background(), bson.M{"data": data})
	if err != nil {
		return nil, nil
	}
	fmt.Printf("function Save | data to MongoDB %s", data)
	return nil, nil
}

func (m *MongoDB) CreateUserStorage(userModel *domain.User) (string, error) {
	_, err := m.Collection.InsertOne(context.Background(), bson.M{"data": userModel})
	if err != nil {
		return "", err
	}

	return "User Created", nil
}
