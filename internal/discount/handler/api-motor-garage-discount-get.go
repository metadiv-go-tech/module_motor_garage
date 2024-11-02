package handler

import (
	"errors"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/discount/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageDiscountGet = metagin.Get(
	"getDiscount",
	"Get Discount",
	"/motor-garage/discount/:id",
	func(ctx metagin.Context[base.RequestPathId, dto.MotorGarageDiscount]) {

		j := ctx.Jwt()

		d := repo.MotorGarageDiscountRepo.FindById(ctx.DB(), ctx.Request().ID, j.GetWorkspaceId())
		if d == nil {
			ctx.Err(errors.New("discount not found"))
			return
		}

		ctx.OK(d.ToDTO())
	},
)
