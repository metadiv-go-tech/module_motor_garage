package request

type MotorGarageInvoiceService struct {
	ID        uint `json:"id"`
	ServiceId uint `json:"service_id"`
	MotorGarageServiceCreate
}
