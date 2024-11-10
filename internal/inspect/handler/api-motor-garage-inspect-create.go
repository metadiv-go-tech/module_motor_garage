package handler

import (
	"errors"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/inspect/repo"
	invoiceRepo "github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

var ApiMotorGarageInspectCreate = metagin.Post(
	"createInspect",
	"Create motor inspect",
	"/motor-garage/inspect",
	func(ctx metagin.Context[request.MotorGarageInspectCreate, dto.MotorGarageInspect]) {

		if err := ctx.Request().Validate(); err != nil {
			ctx.Err(err)
			return
		}

		inspect := ctx.Request().ToEntity(nil)

		invoice := invoiceRepo.MotorGarageInvoiceRepo.FindById(ctx.DB(), inspect.InvoiceId, ctx.WorkspaceId())
		if invoice == nil {
			ctx.Err(errors.New("invoice not found"))
			return
		}
		inspect.Invoice = invoice

		inspect = repo.InspectRepo.Save(ctx.DB(), inspect, ctx.WorkspaceId())
		ctx.OK(inspect.ToDTO(ctx.Locale()))
	},
)
