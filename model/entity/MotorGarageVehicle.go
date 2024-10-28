package entity

import (
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/metaorm"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	relationship "github.com/metadiv-go-tech/module_relationship/model/entity"
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

func (e *MotorGarageVehicle) ToDTO() *dto.MotorGarageVehicle {
	return &dto.MotorGarageVehicle{
		ID:           e.ID,
		Name:         e.Name,
		Year:         e.Year,
		Rego:         metaorm.Encryption.Decrypt(e.Rego),
		RegoExpiry:   e.RegoExpiry,
		Odometer:     e.Odometer,
		VIN:          metaorm.Encryption.Decrypt(e.VIN),
		Registration: metaorm.Encryption.Decrypt(e.Registration),
	}
}
