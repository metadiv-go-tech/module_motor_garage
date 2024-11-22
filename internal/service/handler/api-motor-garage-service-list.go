package handler

import (
	"fmt"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/service/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageServiceList = metagin.Get(
	"listService",
	"List Service",
	fmt.Sprintf("/api/%s/motor-garage/service", config.SystemVersion),
	func(ctx metagin.Context[base.RequestListing, []dto.MotorGarageService]) {

		ps, page := repo.MotorGarageServiceRepo.FindAllComplex(
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
		ds := make([]dto.MotorGarageService, len(ps))
		for i := range ps {
			ds[i] = *ps[i].ToDTO()
		}
		ctx.OK(&ds, page)
	},
)
