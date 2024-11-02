package repo

import (
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

var MotorGarageProductRepo = new(motorGarageProductRepo)

type motorGarageProductRepo struct {
	base.Repository[entity.MotorGarageProduct]
}
