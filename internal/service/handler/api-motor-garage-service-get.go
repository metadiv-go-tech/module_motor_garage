package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/mod_motor_garage_core/repo"
)

func ApiMotorGarageServiceGet(ctx metagin.IContext[base.RequestIDPath]) {
	j := ctx.Jwt()

	e := repo.MotorGarageServiceRepo.FindByID(ctx.GetDB(), ctx.GetRequest().ID, j.GetWorkspaceId())
	if e == nil {
		ctx.Err("service not found")
		return
	}

	ctx.OK(e.ToDTO())
}
