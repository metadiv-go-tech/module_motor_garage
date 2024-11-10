package handler

import (
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metaorm/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

var ApiMotorGarageVehicleList = metagin.Get(
	"listVehicle",
	"List Vehicle",
	"/motor-garage/vehicle",
	func(ctx metagin.Context[request.MotorGarageVehicleList, []dto.MotorGarageVehicle]) {

		bd := metaorm.NewAndQueryBuilder()
		bd.Add(ctx.DB().Or(
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
		))
		if ctx.Request().CustomerId > 0 {
			bd.Add(ctx.DB().Eq("customer_id", ctx.Request().CustomerId))
		}

		ps, page := repo.MotorGarageVehicleRepo.FindAllComplex(
			ctx.DB().Preload("Customer"),
			bd.Build(),
			ctx.Page(),
			ctx.Sort(),
			ctx.WorkspaceId(),
		)

		ds := make([]dto.MotorGarageVehicle, len(ps))
		for i := range ps {
			ds[i] = *ps[i].ToDTO(ctx.Locale())
		}

		ctx.OK(&ds, page)
	},
)
