package handler

import (
	"errors"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageVehicleGet = metagin.Get(
	"getVehicle",
	"Get Vehicle",
	"/motor-garage/vehicle/:id",
	func(ctx metagin.Context[base.RequestPathId, dto.MotorGarageVehicle]) {
		j := ctx.Jwt()

		v := repo.MotorGarageVehicleRepo.FindById(ctx.DB().Preload("Customer"), ctx.Request().ID, j.GetWorkspaceId())
		if v == nil {
			ctx.Err(errors.New("vehicle not found"))
			return
		}

		ctx.OK(v.ToDTO(ctx.Locale()))
	},
)
