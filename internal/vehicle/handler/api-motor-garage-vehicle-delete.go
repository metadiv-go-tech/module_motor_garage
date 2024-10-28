package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
)

func ApiMotorGarageVehicleDelete(ctx metagin.IContext[base.RequestIDPath]) {
	j := ctx.Jwt()

	e := repo.MotorGarageVehicleRepo.FindByID(ctx.GetDB(), ctx.GetRequest().ID, j.GetWorkspaceId())
	if e == nil {
		ctx.Err("vehicle not found")
		return
	}

	if !repo.MotorGarageVehicleRepo.Delete(ctx.GetDB(), e) {
		ctx.InternalServerError("failed to delete vehicle")
		return
	}
	ctx.OK(nil)
}
