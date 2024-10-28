package entity

import "github.com/metadiv-go-tech/metagin/base"

type MotorGarageDiscount struct {
	base.Model
	base.ModelWorkspace

	Name               []byte  `json:"name"`
	Description        []byte  `json:"description"`
	DiscountAmount     uint    `json:"discount_amount"`
	DiscountPercentage float64 `json:"discount_percentage"`
}
