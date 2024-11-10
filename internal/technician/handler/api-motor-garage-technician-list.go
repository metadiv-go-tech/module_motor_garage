package handler

import (
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/technician/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageTechnicianList = metagin.Get(
	"listTechnician",
	"List Technician",
	"/motor-garage/technician",
	func(ctx metagin.Context[base.RequestListing, []dto.MotorGarageTechnician]) {
		technicians, page := repo.TechnicianRepo.FindAllComplex(ctx.DB(), ctx.Request().BuildSimilarClause(
			"name",
		), ctx.Page(), ctx.Sort(), ctx.WorkspaceId())

		dtos := make([]dto.MotorGarageTechnician, len(technicians))
		for i, technician := range technicians {
			dtos[i] = *technician.ToDTO()
		}
		ctx.OK(&dtos, page)
	})
