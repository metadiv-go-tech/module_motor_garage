package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/metadiv-go-tech/gotenberg"
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice_report/service"
	"github.com/metadiv-go-tech/module_motor_garage/internal/util_client"
)

var ApiInvoiceReportBlank = metagin.Get(
	"getInvoiceReportBlank",
	"Get Invoice Report Blank",
	fmt.Sprintf("/api/%s/motor-garage/invoice/report/blank", config.SystemVersion),
	func(ctx metagin.Context[base.Empty, base.Empty]) {
		pdf, err := util_client.HtmlToPdfClient.HtmlToPdf(gotenberg.HTML{
			HTML: service.InvoiceService.GenerateBlankReport(),
		})
		if err != nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New(err.Error()))
			return
		}
		ctx.File(pdf, "invoice-report-blank.pdf")
	},
)
