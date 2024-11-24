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

var ApiStatisticIncomeGet = metagin.Get(
	"getIncome",
	"Income Statistics",
	fmt.Sprintf("/api/%s/motor-garage/statistic/income", config.SystemVersion),
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

		incomeMap := make(map[string]uint)
		for _, invoice := range invoices {
			date := time.Unix(invoice.Date, 0).Format(format)
			incomeMap[date] += invoice.Total
		}

		statistics := make([]dto.MotorGarageStatistic, 0)
		startDate := ctx.Request().From
		for {
			date := time.Unix(startDate, 0).Format(format)
			statistics = append(statistics, dto.MotorGarageStatistic{
				Label: date,
				Value: int(incomeMap[date]),
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

		ctx.OK(&statistics)
	},
)
