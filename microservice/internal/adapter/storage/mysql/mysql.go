package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/fxivan/set_up_server/microservice/configuration"
	mysql_model "github.com/fxivan/set_up_server/microservice/internal/adapter/storage/mysql/model"
	"github.com/fxivan/set_up_server/microservice/internal/core/domain"
	"github.com/fxivan/set_up_server/microservice/internal/core/util"
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

func (m *MySQL) GetUserEmailStorage(userEmail string, collectionName string) (*domain.User, error) {

	//product_id email name password role created_at updated_at
	var userModelMySQL mysql_model.UserModelMySQL
	getUserQuery := `SELECT * FROM users WHERE email=?`
	result, err := m.DB.Query(getUserQuery, userEmail)
	if err != nil {
		return nil, fmt.Errorf("could not create table: %v", err)
	}

	for result.Next() {
		err := result.Scan(&userModelMySQL.ID, &userModelMySQL.Email, &userModelMySQL.Name, &userModelMySQL.Password, &userModelMySQL.Role, &userModelMySQL.CreatedAt, &userModelMySQL.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("could not scan user: %v", err)
		}

	}
	fmt.Print(userModelMySQL)

	//createdAtParse, err := time.Parse("2006-01-02 15:04:05", userModelMySQL.CreatedAt)
	createdAtParse, err := util.ConvertTimeMySQLToTimeTime("2006-01-02 15:04:05", userModelMySQL.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAtParse, err := util.ConvertTimeMySQLToTimeTime("2006-01-02 15:04:05", userModelMySQL.UpdatedAt)
	if err != nil {
		return nil, err
	}

	modelUser := &domain.User{
		ID:        userModelMySQL.ID,
		Email:     userModelMySQL.Email,
		Name:      userModelMySQL.Name,
		Password:  userModelMySQL.Password,
		Role:      domain.UserRole(userModelMySQL.Role),
		CreatedAt: createdAtParse,
		UpdatedAt: updatedAtParse,
	}

	fmt.Printf("Function GetUserEmailStorage | data from MySQL %s", modelUser)
	return modelUser, nil
}

func (m *MySQL) ListUsersStorage(collectionName string) ([]domain.User, error) {

	getUserQuery := `SELECT * FROM users`
	result, err := m.DB.Query(getUserQuery)
	if err != nil {
		return nil, fmt.Errorf("could not create table: %v", err)
	}

	defer result.Close()

	allUsers := make([]domain.User, 0)
	for result.Next() { //result.Next() iteracion por cada fila que arroja el "result" de la consulta
		newUsersModel := new(mysql_model.UserModelMySQL)
		err := result.Scan(&newUsersModel.ID, &newUsersModel.Email, &newUsersModel.Name, &newUsersModel.Password, &newUsersModel.Role, &newUsersModel.CreatedAt, &newUsersModel.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("could not scan user: %v", err)
		}

		createdAtParse, err := util.ConvertTimeMySQLToTimeTime("2006-01-02 15:04:05", newUsersModel.CreatedAt)
		if err != nil {
			return nil, err
		}
		updatedAtParse, err := util.ConvertTimeMySQLToTimeTime("2006-01-02 15:04:05", newUsersModel.UpdatedAt)
		if err != nil {
			return nil, err
		}

		allUsers = append(allUsers, domain.User{
			ID:        newUsersModel.ID,
			Email:     newUsersModel.Email,
			Name:      newUsersModel.Name,
			Password:  newUsersModel.Password,
			Role:      domain.UserRole(newUsersModel.Role),
			CreatedAt: createdAtParse,
			UpdatedAt: updatedAtParse,
		},
		)
	}

	return allUsers, nil
}

func (m *MySQL) GetUserStorage(idUser string, collectionName string) (*domain.User, error) {
	return nil, nil
}
