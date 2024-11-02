package module_motor_garage

import (
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/service/handler"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

func SetupService() {
	metagin.Migrate(&entity.MotorGarageService{})
	metagin.RegisterHandler(
		handler.ApiMotorGarageServiceCreate,
		handler.ApiMotorGarageServiceList,
		handler.ApiMotorGarageServiceGet,
		handler.ApiMotorGarageServiceUpdate,
		handler.ApiMotorGarageServiceDelete,
	)
}
