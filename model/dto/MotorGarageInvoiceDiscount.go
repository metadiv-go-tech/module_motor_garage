package dto

type MotorGarageInvoiceDiscount struct {
	MotorGarageDiscount
	DiscountId *uint `json:"discount_id"`
}
