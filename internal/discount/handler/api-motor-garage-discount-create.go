package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/module_motor_garage/internal/discount/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

func ApiMotorGarageDiscountCreate(ctx metagin.IContext[request.MotorGarageDiscountCreate]) {

	j := ctx.Jwt()

	d := ctx.GetRequest().ToEntity(nil)

	d = repo.MotorGarageDiscountRepo.Save(ctx.GetDB(), d, j.GetWorkspaceId())
	if d == nil {
		ctx.InternalServerError("failed to save discount")
		return
	}

	ctx.OK(d.ToDTO())
}
