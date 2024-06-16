package response

type DataStruct struct {
	OrderID   string  `json:"order_id"`
	Status    string  `json:"status"`
	RefNumber string  `json:"ref_number"`
	Amount    float64 `json:"amount"`
}

type ResVerifyPaymentUala struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    DataStruct `json:"data"`
}
