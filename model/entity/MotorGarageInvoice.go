package entity

import (
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

type MotorGarageInvoice struct {
	base.Model
	base.ModelWorkspace

	Date int64 `json:"date"`

	VehicleId uint                `json:"vehicle_id"`
	Vehicle   *MotorGarageVehicle `json:"vehicle" gorm:"foreignKey:VehicleId"`

	Services  []MotorGarageInvoiceService  `json:"services" gorm:"foreignKey:InvoiceId"`
	Products  []MotorGarageInvoiceProduct  `json:"products" gorm:"foreignKey:InvoiceId"`
	Discounts []MotorGarageInvoiceDiscount `json:"discounts" gorm:"foreignKey:InvoiceId"`
}

func (e *MotorGarageInvoice) ToDTO(locale string) *dto.MotorGarageInvoice {
	d := &dto.MotorGarageInvoice{
		ID:        e.ID,
		Date:      e.Date,
		VehicleId: e.VehicleId,
	}
	if e.Vehicle != nil {
		d.Vehicle = e.Vehicle.ToDTO(locale)
	}
	var total uint
	d.Services = make([]dto.MotorGarageService, len(e.Services))
	for i, s := range e.Services {
		d.Services[i] = *s.ToDTO()
		total += s.PriceAfterTax
	}
	d.Products = make([]dto.MotorGarageInvoiceProduct, len(e.Products))
	for i, p := range e.Products {
		d.Products[i] = *p.ToDTO()
		total += p.PriceAfterTax * p.Quantity
	}
	d.Discounts = make([]dto.MotorGarageDiscount, len(e.Discounts))
	for i, discount := range e.Discounts {
		d.Discounts[i] = *discount.ToDTO()
		total -= discount.DiscountAmount
		total = uint(float64(total) * (1 - discount.DiscountPercentage))
	}
	d.Total = total
	return d
}
