package handler

import (
	"errors"

	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/inspect/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageInspectGet = metagin.Get(
	"getInspect",
	"Get motor inspect",
	"/motor-garage/inspect/:id",
	func(ctx metagin.Context[base.RequestIDPath, dto.MotorGarageInspect]) {
		inspect := repo.InspectRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if inspect == nil {
			ctx.Err(errors.New("inspect not found"))
			return
		}
		ctx.OK(inspect.ToDTO(ctx.Locale()))
	},
)
