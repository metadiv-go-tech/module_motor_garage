package entity

import "github.com/metadiv-go-tech/module_motor_garage/model/dto"

type MotorGarageInvoiceProduct struct {
	MotorGarageProduct

	Quantity uint `json:"quantity"`

	InvoiceId uint                `json:"invoice_id"`
	Invoice   *MotorGarageInvoice `json:"invoice" gorm:"references:ID"`

	ProductId *uint               `json:"product_id"`
	Product   *MotorGarageProduct `json:"product" gorm:"foreignKey:ProductId"`
}

func (e *MotorGarageInvoiceProduct) ToDTO() *dto.MotorGarageInvoiceProduct {
	d := &dto.MotorGarageInvoiceProduct{
		MotorGarageProduct: *e.MotorGarageProduct.ToDTO(),
		Quantity:           e.Quantity,
		ProductId:          e.ProductId,
	}
	return d
}
