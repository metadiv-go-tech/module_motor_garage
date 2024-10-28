package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
)

func ApiMotorGarageInvoiceDelete(ctx metagin.IContext[base.RequestIDPath]) {
	j := ctx.Jwt()

	e := repo.MotorGarageInvoiceRepo.FindByID(ctx.GetDB(), ctx.GetRequest().ID, j.GetWorkspaceId())
	if e == nil {
		ctx.Err("invoice not found")
		return
	}

	if !repo.MotorGarageInvoiceRepo.Delete(ctx.GetDB(), e) {
		ctx.InternalServerError("failed to delete invoice")
		return
	}

	ctx.OK(nil)
}
