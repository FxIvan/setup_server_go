package request

type RegisterUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateGiftCardRequest struct {
	Title         string  `json:"title" binding:"required"`
	Email         string  `json:"email" binding:"required"`
	Description   string  `json:"description" binding:"required"`
	AmountCoupons int     `json:"amount" binding:"required"`
	PriceCoupons  float64 `json:"price" binding:"required"`
}

type InsertCodeRequest struct {
	Code   string `json:"code" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Mobile string `json:"mobile"`
	Name   string `json:"name"`
	CVU    string `json:"cvu"`
	Alias  string `json:"alias"`
	Wallet string `json:"wallet"`
	Red    string `json:"red"`
}
