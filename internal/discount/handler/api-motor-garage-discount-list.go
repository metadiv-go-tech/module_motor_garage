package handler

import (
	"fmt"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/discount/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageDiscountList = metagin.Get(
	"listDiscounts",
	"List Discounts",
	fmt.Sprintf("/api/%s/motor-garage/discount", config.SystemVersion),
	func(ctx metagin.Context[base.RequestListing, []dto.MotorGarageDiscount]) {

		ps, page := repo.MotorGarageDiscountRepo.FindAllComplex(
			ctx.DB(),
			ctx.DB().Or(
				ctx.Request().BuildSimilarClause(
					"name",
					"description",
				),
			),
			ctx.Page(),
			ctx.Sort(),
			ctx.WorkspaceId(),
		)

		ds := make([]dto.MotorGarageDiscount, len(ps))
		for i := range ps {
			ds[i] = *ps[i].ToDTO()
		}

		ctx.OK(&ds, page)
	},
)
