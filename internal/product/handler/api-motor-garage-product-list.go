package handler

import (
	"fmt"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/product/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageProductList = metagin.Get(
	"listProduct",
	"List Product",
	fmt.Sprintf("/api/%s/motor-garage/product", config.SystemVersion),
	func(ctx metagin.Context[base.RequestListing, []dto.MotorGarageProduct]) {

		ps, page := repo.MotorGarageProductRepo.FindAllComplex(
			ctx.DB(),
			ctx.DB().Or(
				ctx.Request().BuildDecryptedSimilarClause(
					"name",
					"description",
				),
				ctx.Request().BuildSimilarClause(
					"price",
					"price_after_tax",
				),
			),
			ctx.Page(),
			ctx.Sort(),
			ctx.WorkspaceId(),
		)
		ds := make([]dto.MotorGarageProduct, len(ps))
		for i := range ps {
			ds[i] = *ps[i].ToDTO()
		}
		ctx.OK(&ds, page)
	},
)
