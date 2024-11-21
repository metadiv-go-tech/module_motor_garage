package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
)

var ApiMotorGarageInvoiceDelete = metagin.Delete(
	"deleteInvoice",
	"Delete Invoice",
	fmt.Sprintf("/api/%s/motor-garage/invoice/:id", config.SystemVersion),
	func(ctx metagin.Context[base.RequestPathId, base.Empty]) {

		e := repo.MotorGarageInvoiceRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if e == nil {
			ctx.Err(errors.New("invoice not found"))
			return
		}

		if !repo.MotorGarageInvoiceRepo.Delete(ctx.DB(), e) {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to delete invoice"))
			return
		}

		ctx.OK(nil)
	},
)
