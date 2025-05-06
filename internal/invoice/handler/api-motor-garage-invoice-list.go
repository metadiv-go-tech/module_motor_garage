package handler

import (
	"fmt"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metaorm/v2"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

var ApiMotorGarageInvoiceList = metagin.Get(
	"listInvoices",
	"List Invoices",
	fmt.Sprintf("/api/%s/motor-garage/invoice", config.SystemVersion),
	func(ctx metagin.Context[request.MotorGarageInvoiceListing, []dto.MotorGarageInvoice]) {

		b := metaorm.NewAndQueryBuilder()
		b.Add(ctx.DB().Or(
			ctx.Request().BuildSimilarClause(
				"Vehicle.name",
				"Vehicle.year",
				"Vehicle__Customer.code",
			),
			ctx.Request().BuildDecryptedSimilarClause(
				"Vehicle.rego",
				"Vehicle.vin",
				"Vehicle.registration",
				"Vehicle__Customer.display_name",
				"Vehicle__Customer.company_name",
				"Vehicle__Customer.search_full_name",
			),
		))

		if ctx.Request().From > 0 {
			b.Add(ctx.DB().Gte("date", ctx.Request().From))
		}
		if ctx.Request().To > 0 {
			b.Add(ctx.DB().Lte("date", ctx.Request().To))
		}

		is, page := repo.MotorGarageInvoiceRepo.FindAllComplexJoined(
			ctx.DB().Joins("Vehicle", "Vehicle.Customer").Preload("Services", "Products", "Discounts"),
			b.Build(),
			"motor_garage_invoices",
			ctx.Page(),
			ctx.Sort(),
			ctx.WorkspaceId(),
		)

		ds := make([]dto.MotorGarageInvoice, 0)
		for i := range is {
			ds = append(ds, *is[i].ToDTO(ctx.Locale()))
		}

		ctx.OK(&ds, page)
	},
)
