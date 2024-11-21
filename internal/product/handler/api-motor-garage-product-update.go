package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/product/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

var ApiMotorGarageProductUpdate = metagin.Put(
	"updateProduct",
	"Update Product",
	fmt.Sprintf("/api/%s/motor-garage/product/:id", config.SystemVersion),
	func(ctx metagin.Context[request.MotorGarageProductUpdate, dto.MotorGarageProduct]) {

		if errMsg := ctx.Request().Validate(); errMsg != "" {
			ctx.Err(errors.New(errMsg))
			return
		}

		product := repo.MotorGarageProductRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if product == nil {
			ctx.Err(errors.New("product not found"))
			return
		}

		product = ctx.Request().ToEntity(product)
		product = repo.MotorGarageProductRepo.Save(ctx.DB(), product, ctx.WorkspaceId())
		if product == nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save product"))
			return
		}

		ctx.OK(product.ToDTO())
	},
)
