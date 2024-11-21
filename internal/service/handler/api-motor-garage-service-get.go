package handler

import (
	"errors"
	"fmt"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/service/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageServiceGet = metagin.Get(
	"getService",
	"Get Service",
	fmt.Sprintf("/api/%s/motor-garage/service/:id", config.SystemVersion),
	func(ctx metagin.Context[base.RequestPathId, dto.MotorGarageService]) {

		e := repo.MotorGarageServiceRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if e == nil {
			ctx.Err(errors.New("service not found"))
			return
		}

		ctx.OK(e.ToDTO())
	},
)
