package handler

import (
	"errors"
	"fmt"

	"github.com/metadiv-go-tech/gotenberg"
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
	"github.com/metadiv-go-tech/module_motor_garage/internal/test_report/service"
	"github.com/metadiv-go-tech/module_motor_garage/internal/util_client"
)

var ApiTestReportPrint = metagin.Get(
	"printTestReport",
	"Print Test Report",
	"/motor-garage/test-report/:id",
	func(ctx metagin.Context[base.RequestPathId, base.Empty]) {

		invoice := repo.MotorGarageInvoiceRepo.FindById(
			ctx.DB().Preload("Vehicle", "Vehicle.Customer", "Vehicle.Customer.ContactPerson", "Services", "Discounts", "Products"),
			ctx.Request().ID, ctx.WorkspaceId())
		if invoice == nil {
			ctx.Err(errors.New("invoice not found"))
			return
		}

		pdf, err := util_client.Client.HtmlToPdf(gotenberg.HTML{
			HTML: service.ReportService.GenerateReport(invoice, ctx.Locale()),
		})
		if err != nil {
			ctx.Err(err)
			return
		}
		ctx.File(pdf, fmt.Sprintf("test-report-#%d.pdf", invoice.ID))
	},
)
