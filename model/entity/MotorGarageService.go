package entity

import (
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/metaorm"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

type MotorGarageService struct {
	base.Model
	base.ModelWorkspace

	Name          []byte `json:"name"`
	Description   []byte `json:"description"`
	Price         uint   `json:"price"`
	PriceAfterTax uint   `json:"price_after_tax"`
}

func (e *MotorGarageService) ToDTO() *dto.MotorGarageService {
	d := new(dto.MotorGarageService)
	d.ID = e.ID
	d.Name = metaorm.Encryption.Decrypt(e.Name)
	d.Description = metaorm.Encryption.Decrypt(e.Description)
	return d
}
