package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
)

func ApiMotorGarageInvoiceGet(ctx metagin.IContext[base.RequestIDPath]) {
	j := ctx.Jwt()

	e := repo.MotorGarageInvoiceRepo.FindByID(ctx.GetDB().Preload("Customer").Preload("Items"), ctx.GetRequest().ID, j.GetWorkspaceId())
	if e == nil {
		ctx.Err("invoice not found")
		return
	}
	ctx.OK(e.ToDTO(ctx.Locale()))
}
