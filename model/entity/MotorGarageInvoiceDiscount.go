package entity

type MotorGarageInvoiceDiscount struct {
	MotorGarageDiscount

	InvoiceId uint                `json:"invoice_id"`
	Invoice   *MotorGarageInvoice `json:"invoice" gorm:"foreignKey:InvoiceId;references:ID"`

	DiscountId *uint                `json:"discount_id"`
	Discount   *MotorGarageDiscount `json:"discount" gorm:"foreignKey:DiscountId"`
}
