package entity

import (
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/metaorm/v2"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	relationship "github.com/metadiv-go-tech/module_relationship/v2/model/entity"
)

type MotorGarageBooking struct {
	base.Model
	base.ModelWorkspace

	DateTime int64 `json:"date_time"`

	CustomerId uint                   `json:"customer_id"`
	Customer   *relationship.Customer `json:"customer"`

	VehicleId uint                `json:"vehicle_id"`
	Vehicle   *MotorGarageVehicle `json:"vehicle"`

	Requirement []byte `json:"requirement"`
	Note        []byte `json:"note"`
}

func (m *MotorGarageBooking) ToDTO(locale string) *dto.MotorGarageBooking {
	d := &dto.MotorGarageBooking{
		ID:          m.ID,
		DateTime:    m.DateTime,
		CustomerId:  m.CustomerId,
		VehicleId:   m.VehicleId,
		Requirement: metaorm.Decrypt(m.Requirement),
		Note:        metaorm.Decrypt(m.Note),
	}
	if m.Customer != nil {
		d.Customer = m.Customer.ToDTO(locale)
	}
	if m.Vehicle != nil {
		d.Vehicle = m.Vehicle.ToDTO(locale)
	}
	return d
}
