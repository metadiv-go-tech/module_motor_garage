package dto

import relationship "github.com/metadiv-go-tech/module_relationship/v2/model/dto"

type MotorGarageBooking struct {
	ID       uint  `json:"id"`
	DateTime int64 `json:"date_time"`

	CustomerId uint                   `json:"customer_id"`
	Customer   *relationship.Customer `json:"customer"`

	VehicleId uint                `json:"vehicle_id"`
	Vehicle   *MotorGarageVehicle `json:"vehicle"`

	Requirement string `json:"requirement"`
	Note        string `json:"note"`
}
