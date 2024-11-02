package handler

import (
	"errors"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/service/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageServiceGet = metagin.Get(
	"getService",
	"Get Service",
	"/motor-garage/service/:id",
	func(ctx metagin.Context[base.RequestPathId, dto.MotorGarageService]) {
		j := ctx.Jwt()

		e := repo.MotorGarageServiceRepo.FindById(ctx.DB(), ctx.Request().ID, j.GetWorkspaceId())
		if e == nil {
			ctx.Err(errors.New("service not found"))
			return
		}

		ctx.OK(e.ToDTO())
	},
)
