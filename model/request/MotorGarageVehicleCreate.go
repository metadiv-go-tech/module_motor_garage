package request

import (
	"github.com/metadiv-go-tech/metaorm/v2"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

type MotorGarageVehicleCreate struct {
	Name         string `json:"name"`
	Year         uint   `json:"year"`
	Rego         string `json:"rego"`
	RegoExpiry   int64  `json:"rego_expiry"`
	Odometer     uint   `json:"odometer"`
	VIN          string `json:"vin"`
	Registration string `json:"registration"`
	CustomerId   *uint  `json:"customer_id"`
}

func (r *MotorGarageVehicleCreate) Validate() string {
	if r.Name == "" {
		return "name is required"
	}
	if r.Registration == "" {
		return "registration is required"
	}
	return ""
}

func (r *MotorGarageVehicleCreate) ToEntity(e *entity.MotorGarageVehicle) *entity.MotorGarageVehicle {
	if e == nil {
		e = &entity.MotorGarageVehicle{}
	}
	e.Name = r.Name
	e.Year = r.Year
	e.Rego = metaorm.Encrypt(r.Rego)
	e.RegoExpiry = r.RegoExpiry
	e.Odometer = r.Odometer
	e.VIN = metaorm.Encrypt(r.VIN)
	e.Registration = metaorm.Encrypt(r.Registration)
	e.CustomerId = r.CustomerId
	return e
}
