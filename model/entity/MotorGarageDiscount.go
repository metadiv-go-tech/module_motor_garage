package entity

import (
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/metaorm"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

type MotorGarageDiscount struct {
	base.Model
	base.ModelWorkspace

	Name               []byte  `json:"name"`
	Description        []byte  `json:"description"`
	DiscountAmount     uint    `json:"discount_amount"`
	DiscountPercentage float64 `json:"discount_percentage"`
}

func (e *MotorGarageDiscount) ToDTO() *dto.MotorGarageDiscount {
	return &dto.MotorGarageDiscount{
		ID:                 e.ID,
		Name:               metaorm.Encryption.Decrypt(e.Name),
		Description:        metaorm.Encryption.Decrypt(e.Description),
		DiscountAmount:     e.DiscountAmount,
		DiscountPercentage: e.DiscountPercentage,
	}
}
