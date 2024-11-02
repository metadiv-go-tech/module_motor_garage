package handler

import (
	"errors"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageInvoiceGet = metagin.Get(
	"getInvoice",
	"Get Invoice",
	"/motor-garage/invoice/:id",
	func(ctx metagin.Context[base.RequestPathId, dto.MotorGarageInvoice]) {
		j := ctx.Jwt()

		e := repo.MotorGarageInvoiceRepo.FindById(
			ctx.DB().Preload("Vehicle", "Vehicle.Customer", "Services", "Products", "Discounts"),
			ctx.Request().ID, j.GetWorkspaceId())
		if e == nil {
			ctx.Err(errors.New("invoice not found"))
			return
		}
		ctx.OK(e.ToDTO(ctx.Locale()))
	},
)
