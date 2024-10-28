package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/module_motor_garage/internal/service/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

func ApiMotorGarageServiceUpdate(ctx metagin.IContext[request.MotorGarageServiceUpdate]) {
	j := ctx.Jwt()

	s := repo.MotorGarageServiceRepo.FindByID(ctx.GetDB(), ctx.GetRequest().ID, j.GetWorkspaceId())
	if s == nil {
		ctx.Err("service not found")
		return
	}

	if errMsg := ctx.GetRequest().Validate(); errMsg != "" {
		ctx.Err(errMsg)
		return
	}

	s = ctx.GetRequest().ToEntity(s)
	s = repo.MotorGarageServiceRepo.Save(ctx.GetDB(), s, j.GetWorkspaceId())
	if s == nil {
		ctx.InternalServerError("failed to save service")
		return
	}

	ctx.OK(s.ToDTO())
}
