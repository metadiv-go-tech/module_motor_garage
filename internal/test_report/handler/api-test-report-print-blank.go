package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/metadiv-go-tech/gotenberg"
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/test_report/service"
	"github.com/metadiv-go-tech/module_motor_garage/internal/util_client"
)

var ApiTestReportPrintBlank = metagin.Get(
	"getTestReportPrintBlank",
	"Get Test Report Print Blank",
	fmt.Sprintf("/api/%s/motor-garage/test-report/print/blank", config.SystemVersion),
	func(ctx metagin.Context[base.Empty, base.Empty]) {
		pdf, err := util_client.HtmlToPdfClient.HtmlToPdf(gotenberg.HTML{
			HTML: service.ReportService.GenerateBlankReport(),
		})
		if err != nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New(err.Error()))
			return
		}
		ctx.File(pdf, "test-report-blank.pdf")
	},
)
