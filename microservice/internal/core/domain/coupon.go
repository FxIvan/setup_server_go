package domain

type Coupon struct {
	ID            string
	Owner         string
	Title         string
	Description   string
	AmountCoupons int
	PriceCoupon   int
	Total         int
}

type ResLinks struct {
	CheckoutLink string `json:"checkoutLink"`
	LinkSuccess  string `json:"success"`
	LinkFailed   string `json:"failed"`
}

type ResponseUalabisPOST struct {
	IdTx        string   `json:"id"`
	Type        string   `json:"type"`
	UUID        string   `json:"uuid"`
	OrderNumber string   `json:"orderNumber"`
	Amount      string   `json:"amount"`
	Status      string   `json:"status"`
	RefNumber   string   `json:"refNumber"`
	Links       ResLinks `json:"links"`
}
