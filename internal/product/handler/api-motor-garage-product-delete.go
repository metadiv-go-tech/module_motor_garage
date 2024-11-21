package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/product/repo"
)

var ApiMotorGarageProductDelete = metagin.Delete(
	"deleteProduct",
	"Delete Product",
	fmt.Sprintf("/api/%s/motor-garage/product/:id", config.SystemVersion),
	func(ctx metagin.Context[base.RequestPathId, base.Empty]) {

		product := repo.MotorGarageProductRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if product == nil {
			ctx.Err(errors.New("product not found"))
			return
		}

		if !repo.MotorGarageProductRepo.Delete(ctx.DB(), product) {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to delete product"))
			return
		}

		ctx.OK(nil)
	},
)
