package request

type RequestPaymentMicroservice struct {
	Amount         float64 `json:"amount"`
	Description    string  `json:"description"`
	SuccesResponse string  `json:"succesResponse"`
	FailedResponse string  `json:"failedResponse"`
}
