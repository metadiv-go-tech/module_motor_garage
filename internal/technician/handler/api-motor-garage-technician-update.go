package handler

import (
	"errors"
	"fmt"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/technician/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

var ApiMotorGarageTechnicianUpdate = metagin.Put(
	"updateTechnician",
	"Update Technician",
	fmt.Sprintf("/api/%s/motor-garage/technician/:id", config.SystemVersion),
	func(ctx metagin.Context[request.MotorGarageTechnicianUpdate, dto.MotorGarageTechnician]) {
		if err := ctx.Request().Validate(); err != nil {
			ctx.Err(err)
			return
		}

		technician := repo.TechnicianRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if technician == nil {
			ctx.Err(errors.New("technician not found"))
			return
		}

		technician = ctx.Request().ToEntity(technician)
		technician = repo.TechnicianRepo.Save(ctx.DB(), technician, ctx.WorkspaceId())
		ctx.OK(technician.ToDTO())
	})
