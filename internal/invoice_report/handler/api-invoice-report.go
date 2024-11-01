package handler

import (
	"github.com/metadiv-go-tech/gotenberg"
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/metaorm"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice_report/service"
	"github.com/metadiv-go-tech/module_motor_garage/internal/util_client"
)

func ApiInvoiceReport(ctx metagin.IContext[base.RequestIDPath]) {

	j := ctx.Jwt()

	invoice := repo.MotorGarageInvoiceRepo.FindByID(
		metaorm.Preload(ctx.GetDB(), "Vehicle", "Vehicle.Customer", "Vehicle.Customer.ContactPerson", "Services", "Discounts", "Products"),
		ctx.GetRequest().ID, j.GetWorkspaceId())
	if invoice == nil {
		ctx.Err("invoice not found")
		return
	}

	pdf, err := util_client.Client.HtmlToPdf(gotenberg.HTML{
		HTML: service.InvoiceService.GenerateReport(invoice, ctx.Locale()),
	})
	if err != nil {
		ctx.InternalServerError(err.Error())
		return
	}

	ctx.OKFile(pdf, "invoice.pdf")
}
