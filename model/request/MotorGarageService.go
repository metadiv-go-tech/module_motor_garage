package request

import (
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/metaorm/v2"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

type MotorGarageServiceCreate struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Price         uint   `json:"price"`
	PriceAfterTax uint   `json:"price_after_tax"`
}

type MotorGarageServiceUpdate struct {
	base.RequestPathId
	MotorGarageServiceCreate
}

func (r *MotorGarageServiceCreate) Validate() string {
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

func (r *MotorGarageServiceCreate) ToEntity(e *entity.MotorGarageService) *entity.MotorGarageService {
	if e == nil {
		e = new(entity.MotorGarageService)
	}
	e.Name = metaorm.Encrypt(r.Name)
	e.Description = metaorm.Encrypt(r.Description)
	e.Price = r.Price
	e.PriceAfterTax = r.PriceAfterTax
	return e
}
