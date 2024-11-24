package handler

import (
	"fmt"
	"time"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metaorm/v2"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"

	"github.com/metadiv-go-tech/module_motor_garage/model/dto"

	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

var ApiStatisticNumberOfInvoiceGet = metagin.Get(
	"getNumberOfInvoice",
	"Number of Invoice",
	fmt.Sprintf("/api/%s/motor-garage/statistic/number-of-invoice", config.SystemVersion),
	func(ctx metagin.Context[request.MotorGarageStatisticGet, []dto.MotorGarageStatistic]) {

		b := metaorm.NewAndQueryBuilder()

		if ctx.Request().From > 0 {
			b.Add(ctx.DB().Gte("date", ctx.Request().From))
		}
		if ctx.Request().To > 0 {
			b.Add(ctx.DB().Lte("date", ctx.Request().To))
		}

		invoices := repo.MotorGarageInvoiceRepo.FindAll(ctx.DB(), b.Build(), ctx.WorkspaceId())

		format := "2006-01-02"
		if ctx.Request().Range == "month" {
			format = "2006-01"
		}

		counterMap := make(map[string]int)
		for _, invoice := range invoices {
			date := time.Unix(invoice.Date, 0).Format(format)
			counterMap[date] = counterMap[date] + 1
		}

		counters := make([]dto.MotorGarageStatistic, 0)
		startDate := ctx.Request().From
		for {
			date := time.Unix(startDate, 0).Format(format)
			counters = append(counters, dto.MotorGarageStatistic{
				Label: date,
				Value: counterMap[date],
			})
			if ctx.Request().Range == "month" {
				startDate += 2592000
			} else {
				startDate += 86400
			}
			if startDate > ctx.Request().To {
				break
			}
		}

		ctx.OK(&counters)
	},
)
