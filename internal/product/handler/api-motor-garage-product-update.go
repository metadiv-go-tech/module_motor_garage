package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/mod_motor_garage_core/model/request"
	"github.com/metadiv-go-tech/mod_motor_garage_core/repo"
)

func ApiMotorGarageProductUpdate(ctx metagin.IContext[request.MotorGarageProductUpdate]) {
	j := ctx.Jwt()

	if errMsg := ctx.GetRequest().Validate(); errMsg != "" {
		ctx.Err(errMsg)
		return
	}

	product := repo.MotorGarageProductRepo.FindByID(ctx.GetDB(), ctx.GetRequest().ID, j.GetWorkspaceId())
	if product == nil {
		ctx.Err("product not found")
		return
	}

	product = ctx.GetRequest().ToEntity(product)
	product = repo.MotorGarageProductRepo.Save(ctx.GetDB(), product, j.GetWorkspaceId())
	if product == nil {
		ctx.InternalServerError("failed to save product")
		return
	}

	ctx.OK(product.ToDTO())
}
