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
