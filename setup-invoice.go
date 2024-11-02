package module_motor_garage

import (
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/handler"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

func SetupInvoice() {
	metagin.Migrate(
		&entity.MotorGarageInvoice{},
		&entity.MotorGarageInvoiceService{},
		&entity.MotorGarageInvoiceProduct{},
		&entity.MotorGarageInvoiceDiscount{},
	)
	metagin.RegisterHandler(
		handler.ApiMotorGarageInvoiceCreate,
		handler.ApiMotorGarageInvoiceList,
		handler.ApiMotorGarageInvoiceGet,
		handler.ApiMotorGarageInvoiceUpdate,
		handler.ApiMotorGarageInvoiceDelete,
	)
}
