package module_motor_garage

import (
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice_report/handler"
)

func SetupInvoiceReport() {
	metagin.RegisterHandler(
		handler.ApiInvoiceReport,
	)
}
