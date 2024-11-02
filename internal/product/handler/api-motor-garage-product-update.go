package handler

import (
	"errors"
	"net/http"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/product/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

var ApiMotorGarageProductUpdate = metagin.Put(
	"updateProduct",
	"Update Product",
	"/motor-garage/product/:id",
	func(ctx metagin.Context[request.MotorGarageProductUpdate, dto.MotorGarageProduct]) {
		j := ctx.Jwt()

		if errMsg := ctx.Request().Validate(); errMsg != "" {
			ctx.Err(errors.New(errMsg))
			return
		}

		product := repo.MotorGarageProductRepo.FindById(ctx.DB(), ctx.Request().ID, j.GetWorkspaceId())
		if product == nil {
			ctx.Err(errors.New("product not found"))
			return
		}

		product = ctx.Request().ToEntity(product)
		product = repo.MotorGarageProductRepo.Save(ctx.DB(), product, j.GetWorkspaceId())
		if product == nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save product"))
			return
		}

		ctx.OK(product.ToDTO())
	},
)
