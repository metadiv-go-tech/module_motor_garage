package handler

import (
	"errors"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/technician/repo"
)

var ApiMotorGarageTechnicianDelete = metagin.Delete(
	"deleteTechnician",
	"Delete Technician",
	"/motor-garage/technician/:id",
	func(ctx metagin.Context[base.RequestPathId, any]) {
		technician := repo.TechnicianRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if technician == nil {
			ctx.Err(errors.New("technician not found"))
			return
		}

		repo.TechnicianRepo.Delete(ctx.DB(), technician)
		ctx.OK(nil)
	})
