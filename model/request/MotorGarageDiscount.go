package request

import (
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/metaorm/v2"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

type MotorGarageDiscountCreate struct {
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	DiscountAmount     uint    `json:"discount_amount"`
	DiscountPercentage float64 `json:"discount_percentage"`
}

type MotorGarageDiscountUpdate struct {
	base.RequestPathId
	MotorGarageDiscountCreate
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
	d.Name = metaorm.Encrypt(r.Name)
	d.Description = metaorm.Encrypt(r.Description)
	d.DiscountAmount = r.DiscountAmount
	d.DiscountPercentage = r.DiscountPercentage
	return d
}
