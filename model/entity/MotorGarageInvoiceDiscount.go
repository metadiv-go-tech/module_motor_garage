package entity

type MotorGarageInvoiceDiscount struct {
	MotorGarageDiscount

	InvoiceId uint                `json:"invoice_id"`
	Invoice   *MotorGarageInvoice `json:"invoice"`

	DiscountId *uint                `json:"discount_id"`
	Discount   *MotorGarageDiscount `json:"discount"`
}
