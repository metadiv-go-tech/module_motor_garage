package handler

import (
	"errors"
	"fmt"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/discount/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageDiscountGet = metagin.Get(
	"getDiscount",
	"Get Discount",
	fmt.Sprintf("/api/%s/motor-garage/discount/:id", config.SystemVersion),
	func(ctx metagin.Context[base.RequestPathId, dto.MotorGarageDiscount]) {

		d := repo.MotorGarageDiscountRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if d == nil {
			ctx.Err(errors.New("discount not found"))
			return
		}

		ctx.OK(d.ToDTO())
	},
)
