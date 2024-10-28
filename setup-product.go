package module_motor_garage

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/module_motor_garage/internal/product/handler"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

func SetupProduct() {
	metagin.Migrate(&entity.MotorGarageProduct{})

	metagin.Post("/motor-garage/product", handler.ApiMotorGarageProductCreate, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "createProduct",
			Models:       []any{request.MotorGarageProduct{}, dto.MotorGarageProduct{}},
			Body:         "MotorGarageProductCreate",
			Response:     "MotorGarageProduct",
		},
	})
	metagin.Get("/motor-garage/product", handler.ApiMotorGarageProductList, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "listProducts",
			Models:       []any{dto.MotorGarageProduct{}},
			Forms:        []string{"page", "size", "field", "asc", "keyword"},
			Response:     "MotorGarageProduct[]",
		},
	})
	metagin.Get("/motor-garage/product/:id", handler.ApiMotorGarageProductGet, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "getProduct",
			Models:       []any{dto.MotorGarageProduct{}},
			Paths:        []string{"id"},
			Response:     "MotorGarageProduct",
		},
	})
	metagin.Put("/motor-garage/product/:id", handler.ApiMotorGarageProductUpdate, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "updateProduct",
			Models:       []any{request.MotorGarageProduct{}, dto.MotorGarageProduct{}},
			Paths:        []string{"id"},
			Body:         "MotorGarageProductUpdate",
			Response:     "MotorGarageProduct",
		},
	})
	metagin.Delete("/motor-garage/product/:id", handler.ApiMotorGarageProductDelete, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "deleteProduct",
			Paths:        []string{"id"},
		},
	})
}
