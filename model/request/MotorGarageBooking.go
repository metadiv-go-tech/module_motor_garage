package request

import (
	"errors"

	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/metaorm/v2"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

type MotorGarageBookingCreate struct {
	DateTime int64 `json:"date_time"`

	CustomerId uint `json:"customer_id"`
	VehicleId  uint `json:"vehicle_id"`

	Requirement string `json:"requirement"`
	Note        string `json:"note"`
}

type MotorGarageBookingUpdate struct {
	base.RequestPathId
	MotorGarageBookingCreate
}

func (r *MotorGarageBookingCreate) Validate() error {
	if r.CustomerId == 0 && r.VehicleId == 0 {
		return errors.New("customer or vehicle are required")
	}
	return nil
}

func (r *MotorGarageBookingCreate) ToEntity(e *entity.MotorGarageBooking) *entity.MotorGarageBooking {
	if e == nil {
		e = &entity.MotorGarageBooking{}
	}
	e.DateTime = r.DateTime
	if r.CustomerId != 0 {
		e.CustomerId = &r.CustomerId
	}
	if r.VehicleId != 0 {
		e.VehicleId = &r.VehicleId
	}
	e.Requirement = metaorm.Encrypt(r.Requirement)
	e.Note = metaorm.Encrypt(r.Note)
	return e
}
