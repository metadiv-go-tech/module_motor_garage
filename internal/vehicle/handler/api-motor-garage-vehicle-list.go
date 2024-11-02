package handler

import (
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageVehicleList = metagin.Get(
	"listVehicle",
	"List Vehicle",
	"/motor-garage/vehicle",
	func(ctx metagin.Context[base.RequestListing, []dto.MotorGarageVehicle]) {
		j := ctx.Jwt()

		ps, page := repo.MotorGarageVehicleRepo.FindAllComplex(
			ctx.DB().Preload("Customer"),
			ctx.DB().Or(
				ctx.Request().BuildDecryptedSimilarClause(
					"rego",
					"registration",
					"vin",
				),
				ctx.Request().BuildSimilarClause(
					"name",
					"year",
					"odometer",
				),
			),
			ctx.Page(),
			ctx.Sort(),
			j.GetWorkspaceId(),
		)

		ds := make([]dto.MotorGarageVehicle, len(ps))
		for i := range ps {
			ds[i] = *ps[i].ToDTO(ctx.Locale())
		}

		ctx.OK(&ds, page)
	},
)
