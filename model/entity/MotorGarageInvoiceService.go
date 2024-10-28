package entity

type MotorGarageInvoiceService struct {
	MotorGarageService

	InvoiceId uint                `json:"invoice_id"`
	Invoice   *MotorGarageInvoice `json:"invoice"`

	ServiceId *uint               `json:"service_id"`
	Service   *MotorGarageService `json:"service"`
}
