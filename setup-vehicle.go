package module_motor_garage

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/handler"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

func SetupVehicle() {
	metagin.Migrate(&entity.MotorGarageVehicle{})

	metagin.Get("/vehicle", handler.ApiMotorGarageVehicleList, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "listVehicles",
			Models:       []any{dto.MotorGarageVehicle{}},
			Forms:        []string{"page", "size", "field", "asc", "keyword"},
			Response:     "MotorGarageVehicle[]",
		},
	})
	metagin.Get("/vehicle/:id", handler.ApiMotorGarageVehicleGet, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "getVehicle",
			Models:       []any{dto.MotorGarageVehicle{}},
			Paths:        []string{"id"},
			Response:     "MotorGarageVehicle",
		},
	})
	metagin.Post("/vehicle", handler.ApiMotorGarageVehicleCreate, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "createVehicle",
			Models:       []any{request.MotorGarageVehicleCreate{}, dto.MotorGarageVehicle{}},
			Body:         "MotorGarageVehicleCreate",
			Response:     "MotorGarageVehicle",
		},
	})
	metagin.Put("/vehicle/:id", handler.ApiMotorGarageVehicleUpdate, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "updateVehicle",
			Models:       []any{request.MotorGarageVehicleUpdate{}, dto.MotorGarageVehicle{}},
			Paths:        []string{"id"},
			Body:         "MotorGarageVehicleUpdate",
			Response:     "MotorGarageVehicle",
		},
	})
	metagin.Delete("/vehicle/:id", handler.ApiMotorGarageVehicleDelete, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "deleteVehicle",
			Paths:        []string{"id"},
		},
	})
}
