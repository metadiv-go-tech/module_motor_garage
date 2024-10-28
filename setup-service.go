package module_motor_garage

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/module_motor_garage/internal/service/handler"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

func SetupService() {
	metagin.Migrate(&entity.MotorGarageService{})

	metagin.Post("/motor-garage/service", handler.ApiMotorGarageServiceCreate, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "createService",
			Models:       []any{request.MotorGarageService{}, dto.MotorGarageService{}},
			Body:         "MotorGarageServiceCreate",
			Response:     "MotorGarageService",
		},
	})
	metagin.Get("/motor-garage/service", handler.ApiMotorGarageServiceList, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "listServices",
			Models:       []any{dto.MotorGarageService{}},
			Forms:        []string{"page", "size", "field", "asc", "keyword"},
			Response:     "MotorGarageService[]",
		},
	})
	metagin.Get("/motor-garage/service/:id", handler.ApiMotorGarageServiceGet, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "getService",
			Models:       []any{dto.MotorGarageService{}},
			Paths:        []string{"id"},
			Response:     "MotorGarageService",
		},
	})
	metagin.Put("/motor-garage/service/:id", handler.ApiMotorGarageServiceUpdate, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "updateService",
			Models:       []any{request.MotorGarageService{}, dto.MotorGarageService{}},
			Paths:        []string{"id"},
			Body:         "MotorGarageServiceUpdate",
			Response:     "MotorGarageService",
		},
	})
	metagin.Delete("/motor-garage/service/:id", handler.ApiMotorGarageServiceDelete, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "deleteService",
			Paths:        []string{"id"},
		},
	})
}
