package module_motor_garage

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice_report/handler"
)

func SetupInvoiceReport() {
	metagin.Get("/motor-garage/invoice/report/:id", handler.ApiInvoiceReport)
}
