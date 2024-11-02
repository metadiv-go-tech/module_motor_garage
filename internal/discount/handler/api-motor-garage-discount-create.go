package handler

import (
	"errors"
	"net/http"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/discount/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

var ApiMotorGarageDiscountCreate = metagin.Post(
	"createDiscount",
	"Create Discount",
	"/motor-garage/discount",
	func(ctx metagin.Context[request.MotorGarageDiscountCreate, dto.MotorGarageDiscount]) {

		j := ctx.Jwt()

		d := ctx.Request().ToEntity(nil)

		d = repo.MotorGarageDiscountRepo.Save(ctx.DB(), d, j.GetWorkspaceId())
		if d == nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save discount"))
			return
		}

		ctx.OK(d.ToDTO())
	},
)
