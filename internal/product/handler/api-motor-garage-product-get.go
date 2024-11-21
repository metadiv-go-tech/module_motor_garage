package handler

import (
	"errors"
	"fmt"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/product/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageProductGet = metagin.Get(
	"getProduct",
	"Get Product",
	fmt.Sprintf("/api/%s/motor-garage/product/:id", config.SystemVersion),
	func(ctx metagin.Context[base.RequestPathId, dto.MotorGarageProduct]) {

		product := repo.MotorGarageProductRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if product == nil {
			ctx.Err(errors.New("product not found"))
			return
		}

		ctx.OK(product.ToDTO())
	},
)
