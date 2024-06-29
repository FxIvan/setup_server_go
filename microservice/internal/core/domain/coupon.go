package domain

type Coupon struct {
	IDReference   string
	Email         string
	Owner         string
	Title         string
	Description   string
	AmountCoupons int
	PriceCoupon   float64
	Total         float64
}

type ResLinks struct {
	CheckoutLink string `json:"checkoutLink"`
	LinkSuccess  string `json:"success"`
	LinkFailed   string `json:"failed"`
}

type ResponseUalabisPOST struct {
	IDReference string
	IdTx        string   `json:"id"`
	Type        string   `json:"type"`
	UUID        string   `json:"uuid"`
	OrderNumber string   `json:"orderNumber"`
	Amount      string   `json:"amount"`
	Status      string   `json:"status"`
	RefNumber   string   `json:"refNumber"`
	Links       ResLinks `json:"links"`
}

type ResponseUalabisPOSTVerify struct {
	OrderID   string `json:"order_id"`
	Status    string `json:"status"`
	RefNumber string `json:"ref_number"`
	Amount    int    `json:"amount"`
}
