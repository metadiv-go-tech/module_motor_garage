package module_motor_garage

import (
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/handler"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

func SetupVehicle() {
	metagin.Migrate(&entity.MotorGarageVehicle{})

	metagin.RegisterHandler(
		handler.ApiMotorGarageVehicleList,
		handler.ApiMotorGarageVehicleGet,
		handler.ApiMotorGarageVehicleCreate,
		handler.ApiMotorGarageVehicleUpdate,
		handler.ApiMotorGarageVehicleDelete,
	)
}
