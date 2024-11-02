package handler

import (
	"github.com/metadiv-go-tech/gotenberg"
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/test_report/template"
	"github.com/metadiv-go-tech/module_motor_garage/internal/util_client"
)

var ApiTestReportPrint = metagin.Get(
	"printTestReport",
	"Print Test Report",
	"/motor-garage/test-report/print",
	func(ctx metagin.Context[base.Empty, base.Empty]) {
		pdf, err := util_client.Client.HtmlToPdf(gotenberg.HTML{
			HTML: template.TestReportTemplate,
		})
		if err != nil {
			ctx.Err(err)
			return
		}
		ctx.File(pdf, "test-report.pdf")
	},
)
