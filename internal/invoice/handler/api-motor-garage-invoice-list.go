package handler

import (
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/metaorm/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageInvoiceList = metagin.Get(
	"listInvoices",
	"List Invoices",
	"/motor-garage/invoice",
	func(ctx metagin.Context[base.RequestListing, []dto.MotorGarageInvoice]) {
		j := ctx.Jwt()

		b := metaorm.NewAndQueryBuilder()
		b.Add(ctx.DB().Or(
			ctx.Request().BuildSimilarClause(
				"date",
			),
		))

		is, page := repo.MotorGarageInvoiceRepo.FindAllComplex(
			ctx.DB().Preload("Vehicle", "Vehicle.Customer", "Services", "Products", "Discounts"),
			b.Build(),
			ctx.Page(),
			ctx.Sort(),
			j.GetWorkspaceId(),
		)

		ds := make([]dto.MotorGarageInvoice, 0)
		for i := range is {
			ds = append(ds, *is[i].ToDTO(ctx.Locale()))
		}

		ctx.OK(&ds, page)
	},
)
