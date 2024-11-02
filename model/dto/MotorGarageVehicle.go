package dto

import (
	relationship "github.com/metadiv-go-tech/module_relationship/v2/model/dto"
)

type MotorGarageVehicle struct {
	ID uint `json:"id"`

	Name         string `json:"name"`
	Year         uint   `json:"year"`
	Rego         string `json:"rego"`
	RegoExpiry   int64  `json:"rego_expiry"`
	Odometer     uint   `json:"odometer"`
	VIN          string `json:"vin"`
	Registration string `json:"registration"`

	CustomerId *uint                  `json:"customer_id"`
	Customer   *relationship.Customer `json:"customer"`
}
