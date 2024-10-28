package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/module_motor_garage/internal/product/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

func ApiMotorGarageProductCreate(ctx metagin.IContext[request.MotorGarageProductCreate]) {

	j := ctx.Jwt()

	if errMsg := ctx.GetRequest().Validate(); errMsg != "" {
		ctx.Err(errMsg)
		return
	}

	product := ctx.GetRequest().ToEntity(nil)

	product = repo.MotorGarageProductRepo.Save(ctx.GetDB(), product, j.GetWorkspaceId())
	if product == nil {
		ctx.InternalServerError("failed to save product")
		return
	}

	ctx.OK(product.ToDTO())

}
