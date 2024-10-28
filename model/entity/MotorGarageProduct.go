package entity

import (
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/metaorm"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

type MotorGarageProduct struct {
	base.Model
	base.ModelWorkspace

	Name          []byte `json:"name"`
	Description   []byte `json:"description"`
	Price         uint   `json:"price"`
	PriceAfterTax uint   `json:"price_after_tax"`
}

func (e *MotorGarageProduct) ToDTO() *dto.MotorGarageProduct {
	d := new(dto.MotorGarageProduct)
	d.ID = e.ID
	d.Name = metaorm.Encryption.Decrypt(e.Name)
	d.Description = metaorm.Encryption.Decrypt(e.Description)
	d.Price = e.Price
	d.PriceAfterTax = e.PriceAfterTax
	return d
}
