package request

import (
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/metaorm"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

type MotorGarageService struct {
	base.RequestIDPath
	Name          string `json:"name"`
	Description   string `json:"description"`
	Price         uint   `json:"price"`
	PriceAfterTax uint   `json:"price_after_tax"`
}

func (r *MotorGarageService) Validate() string {
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

func (r *MotorGarageService) ToEntity(e *entity.MotorGarageService) *entity.MotorGarageService {
	if e == nil {
		e = new(entity.MotorGarageService)
	}
	e.Name = metaorm.Encryption.Encrypt(r.Name)
	e.Description = metaorm.Encryption.Encrypt(r.Description)
	e.Price = r.Price
	e.PriceAfterTax = r.PriceAfterTax
	return e
}
