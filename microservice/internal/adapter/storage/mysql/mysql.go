package mysql

import (
	"database/sql"
	"fmt"

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

func (m *MySQL) CreateUserStorage(userModel *domain.User, collectionName string) (string, error) {
	return "", nil
}
