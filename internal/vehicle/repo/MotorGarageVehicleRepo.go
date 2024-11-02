package repo

import (
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

var MotorGarageVehicleRepo = new(motorGarageVehicleRepo)

type motorGarageVehicleRepo struct {
	base.Repository[entity.MotorGarageVehicle]
}
