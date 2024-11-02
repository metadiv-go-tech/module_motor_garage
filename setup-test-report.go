package module_motor_garage

import (
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/test_report/handler"
)

func SetupTestReport() {
	metagin.RegisterHandler(
		handler.ApiTestReportPrint,
	)
}
