package repo

import (
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

var MotorGarageServiceRepo = new(motorGarageServiceRepo)

type motorGarageServiceRepo struct {
	base.Repository[entity.MotorGarageService]
}
