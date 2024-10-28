package dto

type MotorGarageInvoice struct {
	ID uint `json:"id"`

	Date int64 `json:"date"`

	VehicleId uint                `json:"vehicle_id"`
	Vehicle   *MotorGarageVehicle `json:"vehicle"`

	Services  []MotorGarageService        `json:"services"`
	Products  []MotorGarageInvoiceProduct `json:"products"`
	Discounts []MotorGarageDiscount       `json:"discounts"`

	Total uint `json:"total"`
}
