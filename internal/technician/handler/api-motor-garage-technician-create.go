package handler

import (
	"fmt"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/technician/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

var ApiMotorGarageTechnicianCreate = metagin.Post(
	"createTechnician",
	"Create Technician",
	fmt.Sprintf("/api/%s/motor-garage/technician", config.SystemVersion),
	func(ctx metagin.Context[request.MotorGarageTechnicianCreate, dto.MotorGarageTechnician]) {

		if err := ctx.Request().Validate(); err != nil {
			ctx.Err(err)
			return
		}

		technician := ctx.Request().ToEntity(nil)
		technician = repo.TechnicianRepo.Save(ctx.DB(), technician, ctx.WorkspaceId())
		ctx.OK(technician.ToDTO())
	})
