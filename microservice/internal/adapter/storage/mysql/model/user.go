package mysql_model

type UserModelMySQL struct {
	ID        string `json:"product_id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
