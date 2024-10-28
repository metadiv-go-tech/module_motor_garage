package module_motor_garage

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/module_motor_garage/internal/discount/handler"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

func SetupDiscount() {
	metagin.Migrate(&entity.MotorGarageDiscount{})

	metagin.Post("/motor-garage/discount", handler.ApiMotorGarageDiscountCreate, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "createDiscount",
			Models:       []any{request.MotorGarageDiscountCreate{}, dto.MotorGarageDiscount{}},
			Body:         "MotorGarageDiscountCreate",
			Response:     "MotorGarageDiscount",
		},
	})
	metagin.Get("/motor-garage/discount", handler.ApiMotorGarageDiscountList, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "listDiscounts",
			Models:       []any{dto.MotorGarageDiscount{}},
			Forms:        []string{"page", "size", "field", "asc", "keyword"},
			Response:     "MotorGarageDiscount[]",
		},
	})
	metagin.Get("/motor-garage/discount/:id", handler.ApiMotorGarageDiscountGet, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "getDiscount",
			Models:       []any{dto.MotorGarageDiscount{}},
			Paths:        []string{"id"},
			Response:     "MotorGarageDiscount",
		},
	})
	metagin.Put("/motor-garage/discount/:id", handler.ApiMotorGarageDiscountUpdate, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "updateDiscount",
			Models:       []any{request.MotorGarageDiscountUpdate{}, dto.MotorGarageDiscount{}},
			Paths:        []string{"id"},
			Body:         "MotorGarageDiscountUpdate",
			Response:     "MotorGarageDiscount",
		},
	})
	metagin.Delete("/motor-garage/discount/:id", handler.ApiMotorGarageDiscountDelete, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "deleteDiscount",
			Paths:        []string{"id"},
		},
	})
}
