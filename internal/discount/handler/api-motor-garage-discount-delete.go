package handler

import (
	"errors"
	"net/http"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/discount/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageDiscountDelete = metagin.Delete(
	"deleteDiscount",
	"Delete Discount",
	"/motor-garage/discount/:id",
	func(ctx metagin.Context[base.RequestPathId, dto.MotorGarageDiscount]) {

		j := ctx.Jwt()

		i := repo.MotorGarageDiscountRepo.FindById(ctx.DB(), ctx.Request().ID, j.GetWorkspaceId())
		if i == nil {
			ctx.Err(errors.New("discount not found"))
			return
		}

		if !repo.MotorGarageDiscountRepo.Delete(ctx.DB(), i) {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to delete discount"))
			return
		}

		ctx.OK(nil)
	},
)
