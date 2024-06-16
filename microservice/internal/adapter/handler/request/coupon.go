package request

type RequestPaymentMicroservice struct {
	Amount         int    `json:"amount"`
	Description    string `json:"description"`
	SuccesResponse string `json:"succesResponse"`
	FailedResponse string `json:"failedResponse"`
}
