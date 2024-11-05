package dto

type MotorGarageInvoiceService struct {
	MotorGarageService
	ServiceId *uint `json:"service_id"`
}
