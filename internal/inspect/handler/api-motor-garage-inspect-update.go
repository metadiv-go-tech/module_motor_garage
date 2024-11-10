package handler

import (
	"errors"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/inspect/repo"
	invoiceRepo "github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

var ApiMotorGarageInspectUpdate = metagin.Put(
	"updateInspect",
	"Update motor inspect",
	"/motor-garage/inspect/:id",
	func(ctx metagin.Context[request.MotorGarageInspectUpdate, dto.MotorGarageInspect]) {

		if err := ctx.Request().Validate(); err != nil {
			ctx.Err(err)
			return
		}

		inspect := repo.InspectRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if inspect == nil {
			ctx.Err(errors.New("inspect not found"))
			return
		}

		inspect = ctx.Request().ToEntity(inspect)

		invoice := invoiceRepo.MotorGarageInvoiceRepo.FindById(ctx.DB(), inspect.InvoiceId, ctx.WorkspaceId())
		if invoice == nil {
			ctx.Err(errors.New("invoice not found"))
			return
		}
		inspect.Invoice = invoice

		inspect = repo.InspectRepo.Save(ctx.DB(), inspect, ctx.WorkspaceId())
		ctx.OK(inspect.ToDTO(ctx.Locale()))
	})
