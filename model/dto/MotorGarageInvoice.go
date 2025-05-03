package dto

type MotorGarageInvoice struct {
	ID uint `json:"id"`

	Date int64 `json:"date"`

	KM string `json:"km"`

	Vehicle *MotorGarageVehicle `json:"vehicle"`
	Booking *MotorGarageBooking `json:"booking"`
	Inspect *MotorGarageInspect `json:"inspect"`

	Services  []MotorGarageInvoiceService  `json:"services"`
	Products  []MotorGarageInvoiceProduct  `json:"products"`
	Discounts []MotorGarageInvoiceDiscount `json:"discounts"`

	Total uint `json:"total"`
}
