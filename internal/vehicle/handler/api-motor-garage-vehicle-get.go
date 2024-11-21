package handler

import (
	"errors"
	"fmt"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageVehicleGet = metagin.Get(
	"getVehicle",
	"Get Vehicle",
	fmt.Sprintf("/api/%s/motor-garage/vehicle/:id", config.SystemVersion),
	func(ctx metagin.Context[base.RequestPathId, dto.MotorGarageVehicle]) {

		v := repo.MotorGarageVehicleRepo.FindById(ctx.DB().Preload("Customer"), ctx.Request().ID, ctx.WorkspaceId())
		if v == nil {
			ctx.Err(errors.New("vehicle not found"))
			return
		}

		ctx.OK(v.ToDTO(ctx.Locale()))
	},
)
