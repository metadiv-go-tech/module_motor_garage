package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/service/repo"
)

var ApiMotorGarageServiceDelete = metagin.Delete(
	"deleteService",
	"Delete Service",
	fmt.Sprintf("/api/%s/motor-garage/service/:id", config.SystemVersion),
	func(ctx metagin.Context[base.RequestPathId, base.Empty]) {

		e := repo.MotorGarageServiceRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if e == nil {
			ctx.Err(errors.New("service not found"))
			return
		}

		if !repo.MotorGarageServiceRepo.Delete(ctx.DB(), e) {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to delete service"))
			return
		}
		ctx.OK(nil)
	},
)
