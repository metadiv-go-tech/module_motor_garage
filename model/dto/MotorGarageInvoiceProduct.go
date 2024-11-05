package dto

type MotorGarageInvoiceProduct struct {
	MotorGarageProduct
	ProductId *uint `json:"product_id"`
	Quantity  uint  `json:"quantity"`
}
