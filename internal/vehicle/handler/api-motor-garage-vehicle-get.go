package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
)

func ApiMotorGarageVehicleGet(ctx metagin.IContext[base.RequestIDPath]) {
	j := ctx.Jwt()

	v := repo.MotorGarageVehicleRepo.FindByID(ctx.GetDB().Preload("Customer"), ctx.GetRequest().ID, j.GetWorkspaceId())
	if v == nil {
		ctx.Err("vehicle not found")
		return
	}

	ctx.OK(v.ToDTO(ctx.Locale()))
}
