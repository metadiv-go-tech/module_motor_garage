package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/module_motor_garage/internal/service/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

func ApiMotorGarageServiceCreate(ctx metagin.IContext[request.MotorGarageService]) {
	j := ctx.Jwt()

	s := ctx.GetRequest().ToEntity(nil)

	if errMsg := ctx.GetRequest().Validate(); errMsg != "" {
		ctx.Err(errMsg)
		return
	}

	s = repo.MotorGarageServiceRepo.Save(ctx.GetDB(), s, j.GetWorkspaceId())
	if s == nil {
		ctx.InternalServerError("failed to save service")
		return
	}

	ctx.OK(s.ToDTO())
}
