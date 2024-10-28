package request

import (
	"github.com/metadiv-go-tech/metaorm"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

type MotorGarageDiscountCreate struct {
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	DiscountAmount     uint    `json:"discount_amount"`
	DiscountPercentage float64 `json:"discount_percentage"`
}

func (r *MotorGarageDiscountCreate) Validate() string {
	if r.Name == "" {
		return "name is required"
	}
	return ""
}

func (r *MotorGarageDiscountCreate) ToEntity(d *entity.MotorGarageDiscount) *entity.MotorGarageDiscount {
	if d == nil {
		d = &entity.MotorGarageDiscount{}
	}
	d.Name = metaorm.Encryption.Encrypt(r.Name)
	d.Description = metaorm.Encryption.Encrypt(r.Description)
	d.DiscountAmount = r.DiscountAmount
	d.DiscountPercentage = r.DiscountPercentage
	return d
}
