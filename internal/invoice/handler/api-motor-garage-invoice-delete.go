package handler

import (
	"errors"
	"net/http"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
)

var ApiMotorGarageInvoiceDelete = metagin.Delete(
	"deleteInvoice",
	"Delete Invoice",
	"/motor-garage/invoice/:id",
	func(ctx metagin.Context[base.RequestPathId, base.Empty]) {
		j := ctx.Jwt()

		e := repo.MotorGarageInvoiceRepo.FindById(ctx.DB(), ctx.Request().ID, j.GetWorkspaceId())
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
