package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/service/repo"
)

func ApiMotorGarageServiceDelete(ctx metagin.IContext[base.RequestIDPath]) {
	j := ctx.Jwt()

	e := repo.MotorGarageServiceRepo.FindByID(ctx.GetDB(), ctx.GetRequest().ID, j.GetWorkspaceId())
	if e == nil {
		ctx.Err("service not found")
		return
	}

	if !repo.MotorGarageServiceRepo.Delete(ctx.GetDB(), e) {
		ctx.InternalServerError("failed to delete service")
		return
	}
	ctx.OK(nil)
}
