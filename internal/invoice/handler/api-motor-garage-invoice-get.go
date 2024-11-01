package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/metaorm"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
)

func ApiMotorGarageInvoiceGet(ctx metagin.IContext[base.RequestIDPath]) {
	j := ctx.Jwt()

	e := repo.MotorGarageInvoiceRepo.FindByID(
		metaorm.Preload(ctx.GetDB(), "Vehicle", "Vehicle.Customer", "Services", "Products", "Discounts"),
		ctx.GetRequest().ID, j.GetWorkspaceId())
	if e == nil {
		ctx.Err("invoice not found")
		return
	}
	ctx.OK(e.ToDTO(ctx.Locale()))
}
