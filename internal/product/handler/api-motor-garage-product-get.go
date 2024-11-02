package handler

import (
	"errors"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/product/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageProductGet = metagin.Get(
	"getProduct",
	"Get Product",
	"/motor-garage/product/:id",
	func(ctx metagin.Context[base.RequestPathId, dto.MotorGarageProduct]) {
		j := ctx.Jwt()

		product := repo.MotorGarageProductRepo.FindById(ctx.DB(), ctx.Request().ID, j.GetWorkspaceId())
		if product == nil {
			ctx.Err(errors.New("product not found"))
			return
		}

		ctx.OK(product.ToDTO())
	},
)
