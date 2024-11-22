package handler

import (
	"errors"
	"fmt"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/technician/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageTechnicianGet = metagin.Get(
	"getTechnician",
	"Get Technician",
	fmt.Sprintf("/api/%s/motor-garage/technician/:id", config.SystemVersion),
	func(ctx metagin.Context[base.RequestPathId, dto.MotorGarageTechnician]) {
		technician := repo.TechnicianRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if technician == nil {
			ctx.Err(errors.New("technician not found"))
			return
		}
		ctx.OK(technician.ToDTO())
	})
