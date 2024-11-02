package repo

import (
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

var MotorGarageDiscountRepo = new(motorGarageDiscountRepo)

type motorGarageDiscountRepo struct {
	base.Repository[entity.MotorGarageDiscount]
}
