package request

type MotorGarageInvoiceProduct struct {
	ID        uint `json:"id"`
	Quantity  uint `json:"quantity"`
	ProductId uint `json:"product_id"`
	MotorGarageProductCreate
}
