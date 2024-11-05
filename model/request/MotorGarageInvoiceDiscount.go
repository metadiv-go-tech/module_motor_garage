package request

type MotorGarageInvoiceDiscount struct {
	ID         uint `json:"id"`
	DiscountId uint `json:"discount_id"`
	MotorGarageDiscountCreate
}
