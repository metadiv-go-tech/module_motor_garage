package dto

type MotorGarageInvoice struct {
	ID uint `json:"id"`

	Date int64 `json:"date"`

	VehicleId uint                `json:"vehicle_id"`
	Vehicle   *MotorGarageVehicle `json:"vehicle"`

	Services  []MotorGarageInvoiceService  `json:"services"`
	Products  []MotorGarageInvoiceProduct  `json:"products"`
	Discounts []MotorGarageInvoiceDiscount `json:"discounts"`

	Total uint `json:"total"`
}
