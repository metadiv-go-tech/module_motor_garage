package entity

import "github.com/metadiv-go-tech/module_motor_garage/model/dto"

type MotorGarageInvoiceDiscount struct {
	MotorGarageDiscount

	InvoiceId uint                `json:"invoice_id"`
	Invoice   *MotorGarageInvoice `json:"invoice" gorm:"references:ID"`

	DiscountId *uint                `json:"discount_id"`
	Discount   *MotorGarageDiscount `json:"discount" gorm:"foreignKey:DiscountId"`
}

func (e *MotorGarageInvoiceDiscount) ToDTO() *dto.MotorGarageInvoiceDiscount {
	return &dto.MotorGarageInvoiceDiscount{
		MotorGarageDiscount: *e.MotorGarageDiscount.ToDTO(),
		DiscountId:          e.DiscountId,
	}
}
