package entity

import (
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/metaorm/v2"
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
	return &dto.MotorGarageService{
		ID:            e.ID,
		Name:          metaorm.Decrypt(e.Name),
		Description:   metaorm.Decrypt(e.Description),
		Price:         e.Price,
		PriceAfterTax: e.PriceAfterTax,
	}
}
