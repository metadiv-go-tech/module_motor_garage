package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/mod_motor_garage_core/repo"
)

func ApiMotorGarageProductDelete(ctx metagin.IContext[base.RequestIDPath]) {
	j := ctx.Jwt()

	product := repo.MotorGarageProductRepo.FindByID(ctx.GetDB(), ctx.GetRequest().ID, j.GetWorkspaceId())
	if product == nil {
		ctx.Err("product not found")
		return
	}

	if !repo.MotorGarageProductRepo.Delete(ctx.GetDB(), product) {
		ctx.InternalServerError("failed to delete product")
		return
	}

	ctx.OK(nil)
}
