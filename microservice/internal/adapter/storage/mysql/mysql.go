package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/fxivan/set_up_server/microservice/configuration"
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
)

type MySQL struct {
	DB *sql.DB
}

func New(config *configuration.Configuration) (*MySQL, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not connect to mysql: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping mysql: %w", err)
	}

	return &MySQL{DB: db}, nil
}

func (m *MySQL) Find(id string) (*domain.User, error) {
	objExample := &domain.User{
		ID:        "a1sd",
		Name:      "John Doe",
		Email:     "",
		Password:  "password",
		Role:      "admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return objExample, nil

}

func (r *MySQL) FindKey(key string) (*domain.User, error) {
	return nil, nil
}

func (m *MySQL) Save(data string) (*domain.User, error) {
	query := "INSERT INTO my_table (data) VALUES (?)"
	_, err := m.DB.Exec(query, data)
	if err != nil {
		return nil, nil
	}
	fmt.Printf("Function Save | data to mysql %s", data)
	return nil, nil
}

func (m *MySQL) CreateUserStorage(userModel *domain.User) (string, error) {
	return "", nil
}
