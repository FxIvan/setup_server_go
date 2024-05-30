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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, "users")
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
	// Crear la tabla si no existe
	createTableQuery := `
	    CREATE TABLE IF NOT EXISTS users (
	        product_id int primary key auto_increment,
	        email text,
	        name text,
	        password text,
	        role text,
	        created_at datetime default CURRENT_TIMESTAMP,
	        updated_at datetime default CURRENT_TIMESTAMP
	    )
	`
	_, err := m.DB.Exec(createTableQuery)
	if err != nil {
		return "", fmt.Errorf("could not create table: %v", err)
	}

	// Insertar datos del usuario en la tabla
	insertUserQuery := `
	    INSERT INTO users (email, name, password, role) VALUES (?, ?, ?, ?)
	`
	_, err = m.DB.Exec(insertUserQuery, userModel.Email, userModel.Name, userModel.Password, userModel.Role)
	if err != nil {
		return "", fmt.Errorf("could not insert user: %v", err)
	}

	fmt.Printf("Function Save | data to MySQL %s", userModel)
	return "User Created", nil
}
