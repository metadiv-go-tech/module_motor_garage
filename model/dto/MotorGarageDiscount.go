package dto

type MotorGarageDiscount struct {
	ID                 uint    `json:"id"`
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	DiscountAmount     uint    `json:"discount_amount"`
	DiscountPercentage float64 `json:"discount_percentage"`
}
