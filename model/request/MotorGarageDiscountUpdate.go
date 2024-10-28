package request

import "github.com/metadiv-go-tech/metagin/base"

type MotorGarageDiscountUpdate struct {
	base.RequestIDPath
	MotorGarageDiscountCreate
}
