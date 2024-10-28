package entity

import "github.com/metadiv-go-tech/metagin/base"

type MotorGarageService struct {
	base.Model
	base.ModelWorkspace

	Name          []byte `json:"name"`
	Description   []byte `json:"description"`
	Price         uint   `json:"price"`
	PriceAfterTax uint   `json:"price_after_tax"`
}
