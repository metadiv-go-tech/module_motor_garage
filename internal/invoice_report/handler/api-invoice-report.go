package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/metadiv-go-tech/gotenberg"
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice_report/service"
	"github.com/metadiv-go-tech/module_motor_garage/internal/util_client"
)

var ApiInvoiceReport = metagin.Get(
	"getInvoiceReport",
	"Get Invoice Report",
	fmt.Sprintf("/api/%s/motor-garage/invoice/report/:id", config.SystemVersion),
	func(ctx metagin.Context[base.RequestPathId, base.Empty]) {

		invoice := repo.MotorGarageInvoiceRepo.FindById(
			ctx.DB().Preload("Vehicle", "Vehicle.Customer", "Vehicle.Customer.ContactPerson", "Services", "Discounts", "Products"),
			ctx.Request().ID, ctx.WorkspaceId())
		if invoice == nil {
			ctx.Err(errors.New("invoice not found"))
			return
		}

		pdf, err := util_client.HtmlToPdfClient.HtmlToPdf(gotenberg.HTML{
			HTML: service.InvoiceService.GenerateReport(invoice, ctx.Locale()),
		})
		if err != nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New(err.Error()))
			return
		}

		ctx.File(pdf, fmt.Sprintf("invoice-#%d.pdf", invoice.ID))
	},
)
