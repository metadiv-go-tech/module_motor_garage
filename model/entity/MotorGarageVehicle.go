package entity

import (
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/metaorm/v2"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	relationship "github.com/metadiv-go-tech/module_relationship/v2/model/entity"
)

type MotorGarageVehicle struct {
	base.Model
	base.ModelWorkspace

	Name         string `json:"name"`
	Year         uint   `json:"year"`
	Rego         []byte `json:"rego"`
	RegoExpiry   int64  `json:"rego_expiry"`
	Odometer     uint   `json:"odometer"`
	VIN          []byte `json:"vin"`
	Registration []byte `json:"registration"`

	CustomerId *uint                  `json:"customer_id"`
	Customer   *relationship.Customer `json:"customer" gorm:"foreignKey:CustomerId"`
}

func (e *MotorGarageVehicle) ToDTO(locale string) *dto.MotorGarageVehicle {
	d := &dto.MotorGarageVehicle{
		ID:           e.ID,
		Name:         e.Name,
		Year:         e.Year,
		Rego:         metaorm.Decrypt(e.Rego),
		RegoExpiry:   e.RegoExpiry,
		Odometer:     e.Odometer,
		VIN:          metaorm.Decrypt(e.VIN),
		Registration: metaorm.Decrypt(e.Registration),
		CustomerId:   e.CustomerId,
	}
	if e.Customer != nil {
		d.Customer = e.Customer.ToDTO(locale)
	}
	return d
}
