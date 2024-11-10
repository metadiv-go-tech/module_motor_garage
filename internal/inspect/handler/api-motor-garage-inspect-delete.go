package handler

import (
	"errors"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/inspect/repo"
)

var ApiMotorGarageInspectDelete = metagin.Delete(
	"deleteInspect",
	"Delete motor inspect",
	"/motor-garage/inspect/:id",
	func(ctx metagin.Context[base.RequestPathId, base.Empty]) {
		inspect := repo.InspectRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if inspect == nil {
			ctx.Err(errors.New("inspect not found"))
			return
		}
		repo.InspectRepo.Delete(ctx.DB(), inspect)
		ctx.OK(nil)
	},
)
