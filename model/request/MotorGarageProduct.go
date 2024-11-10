package request

import (
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/metaorm/v2"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

type MotorGarageProductCreate struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Price         uint   `json:"price"`
	PriceAfterTax uint   `json:"price_after_tax"`
}

type MotorGarageProductUpdate struct {
	base.RequestPathId
	MotorGarageProductCreate
}

func (r *MotorGarageProductCreate) Validate() string {
	if r.Name == "" {
		return "name is required"
	}
	if r.Price == 0 {
		return "price is required"
	}
	if r.PriceAfterTax == 0 {
		return "price after tax is required"
	}
	return ""
}

func (r *MotorGarageProductCreate) ToEntity(e *entity.MotorGarageProduct) *entity.MotorGarageProduct {
	if e == nil {
		e = new(entity.MotorGarageProduct)
	}
	e.Name = metaorm.Encrypt(r.Name)
	e.Description = metaorm.Encrypt(r.Description)
	e.Price = r.Price
	e.PriceAfterTax = r.PriceAfterTax
	return e
}
