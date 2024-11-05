package entity

import "github.com/metadiv-go-tech/module_motor_garage/model/dto"

type MotorGarageInvoiceService struct {
	MotorGarageService

	InvoiceId uint                `json:"invoice_id"`
	Invoice   *MotorGarageInvoice `json:"invoice" gorm:"references:ID"`

	ServiceId *uint               `json:"service_id"`
	Service   *MotorGarageService `json:"service" gorm:"foreignKey:ServiceId"`
}

func (e *MotorGarageInvoiceService) ToDTO() *dto.MotorGarageInvoiceService {
	return &dto.MotorGarageInvoiceService{
		MotorGarageService: *e.MotorGarageService.ToDTO(),
		ServiceId:          e.ServiceId,
	}
}
