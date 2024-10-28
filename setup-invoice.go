package module_motor_garage

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/handler"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

func SetupInvoice() {
	metagin.Migrate(
		&entity.MotorGarageInvoice{},
		&entity.MotorGarageInvoiceService{},
		&entity.MotorGarageInvoiceProduct{},
		&entity.MotorGarageInvoiceDiscount{},
	)

	metagin.Post("/invoice", handler.ApiMotorGarageInvoiceCreate, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "createInvoice",
			Models:       []any{request.MotorGarageInvoiceCreate{}, dto.MotorGarageInvoice{}},
			Body:         "MotorGarageInvoiceCreate",
			Response:     "MotorGarageInvoice",
		},
	})
	metagin.Get("/invoice", handler.ApiMotorGarageInvoiceList, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "listInvoices",
			Models:       []any{dto.MotorGarageInvoice{}},
			Forms:        []string{"page", "size", "field", "asc", "keyword"},
			Response:     "MotorGarageInvoice[]",
		},
	})
	metagin.Get("/invoice/:id", handler.ApiMotorGarageInvoiceGet, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "getInvoice",
			Models:       []any{dto.MotorGarageInvoice{}},
			Paths:        []string{"id"},
			Response:     "MotorGarageInvoice",
		},
	})
	metagin.Put("/invoice/:id", handler.ApiMotorGarageInvoiceUpdate, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "updateInvoice",
			Models:       []any{request.MotorGarageInvoiceUpdate{}, dto.MotorGarageInvoice{}},
			Paths:        []string{"id"},
			Body:         "MotorGarageInvoiceUpdate",
			Response:     "MotorGarageInvoice",
		},
	})
	metagin.Delete("/invoice/:id", handler.ApiMotorGarageInvoiceDelete, &metagin.ApiHandlerOpts{
		TypescriptOpts: &metagin.TypescriptOpts{
			FunctionName: "deleteInvoice",
			Paths:        []string{"id"},
		},
	})
}
