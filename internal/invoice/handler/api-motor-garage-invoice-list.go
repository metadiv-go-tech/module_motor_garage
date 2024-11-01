package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/metaorm"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

func ApiMotorGarageInvoiceList(ctx metagin.IContext[base.RequestListing]) {
	j := ctx.Jwt()

	cls := make([]metaorm.IClause, 0)
	cls = append(cls, metaorm.Or(
		ctx.GetRequest().BuildSimilarClause(
			"date",
		),
	))

	is, page := repo.MotorGarageInvoiceRepo.FindAllComplex(
		metaorm.Preload(ctx.GetDB(), "Vehicle", "Vehicle.Customer", "Services", "Products", "Discounts"),
		metaorm.And(cls...),
		ctx.Page(),
		ctx.Sort(),
		"",
		j.GetWorkspaceId(),
	)

	ds := make([]dto.MotorGarageInvoice, 0)
	for i := range is {
		ds = append(ds, *is[i].ToDTO(ctx.Locale()))
	}

	ctx.OK(ds, page)
}
