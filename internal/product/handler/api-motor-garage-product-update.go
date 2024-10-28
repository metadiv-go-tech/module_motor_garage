package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/module_motor_garage/internal/product/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
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
