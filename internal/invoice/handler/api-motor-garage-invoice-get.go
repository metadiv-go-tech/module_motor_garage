package handler

import (
	"errors"
	"fmt"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageInvoiceGet = metagin.Get(
	"getInvoice",
	"Get Invoice",
	fmt.Sprintf("/api/%s/motor-garage/invoice/:id", config.SystemVersion),
	func(ctx metagin.Context[base.RequestPathId, dto.MotorGarageInvoice]) {

		e := repo.MotorGarageInvoiceRepo.FindById(
			ctx.DB().Preload("Vehicle", "Vehicle.Customer",
				"Services", "Products", "Discounts", "Booking", "Inspect"),
			ctx.Request().ID, ctx.WorkspaceId())
		if e == nil {
			ctx.Err(errors.New("invoice not found"))
			return
		}
		ctx.OK(e.ToDTO(ctx.Locale()))
	},
)
