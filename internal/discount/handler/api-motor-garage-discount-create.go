package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/discount/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

var ApiMotorGarageDiscountCreate = metagin.Post(
	"createDiscount",
	"Create Discount",
	fmt.Sprintf("/api/%s/motor-garage/discount", config.SystemVersion),
	func(ctx metagin.Context[request.MotorGarageDiscountCreate, dto.MotorGarageDiscount]) {

		d := ctx.Request().ToEntity(nil)

		d = repo.MotorGarageDiscountRepo.Save(ctx.DB(), d, ctx.WorkspaceId())
		if d == nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save discount"))
			return
		}

		ctx.OK(d.ToDTO())
	},
)
