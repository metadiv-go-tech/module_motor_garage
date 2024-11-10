package handler

import (
	"errors"
	"net/http"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/discount/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

var ApiMotorGarageDiscountUpdate = metagin.Put(
	"updateDiscount",
	"Update Discount",
	"/motor-garage/discount",
	func(ctx metagin.Context[request.MotorGarageDiscountUpdate, dto.MotorGarageDiscount]) {

		d := repo.MotorGarageDiscountRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if d == nil {
			ctx.Err(errors.New("discount not found"))
			return
		}

		d = ctx.Request().ToEntity(d)
		d = repo.MotorGarageDiscountRepo.Save(ctx.DB(), d, ctx.WorkspaceId())
		if d == nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save discount"))
			return
		}

		ctx.OK(d.ToDTO())
	},
)
