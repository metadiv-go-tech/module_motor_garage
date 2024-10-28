package dto

import (
	"github.com/metadiv-go-tech/metagin/base"
)

type MotorGarageProduct struct {
	base.Model
	base.ModelWorkspace

	Name          string `json:"name"`
	Description   string `json:"description"`
	Price         uint   `json:"price"`
	PriceAfterTax uint   `json:"price_after_tax"`
}
