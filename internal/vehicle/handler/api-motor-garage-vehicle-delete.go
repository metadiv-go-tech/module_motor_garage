package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
)

var ApiMotorGarageVehicleDelete = metagin.Delete(
	"deleteVehicle",
	"Delete Vehicle",
	fmt.Sprintf("/api/%s/motor-garage/vehicle/:id", config.SystemVersion),
	func(ctx metagin.Context[base.RequestPathId, base.Empty]) {

		e := repo.MotorGarageVehicleRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if e == nil {
			ctx.Err(errors.New("vehicle not found"))
			return
		}

		if !repo.MotorGarageVehicleRepo.Delete(ctx.DB(), e) {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to delete vehicle"))
			return
		}
		ctx.OK(nil)
	},
)
