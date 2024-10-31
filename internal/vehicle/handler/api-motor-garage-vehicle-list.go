package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/metaorm"
	"github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

func ApiMotorGarageVehicleList(ctx metagin.IContext[base.RequestListing]) {
	j := ctx.Jwt()

	ps, page := repo.MotorGarageVehicleRepo.FindAllComplex(
		ctx.GetDB().Preload("Customer"),
		metaorm.Or(
			ctx.GetRequest().BuildDecryptedSimilarClause(
				"rego",
				"registration",
				"vin",
			),
			ctx.GetRequest().BuildSimilarClause(
				"name",
				"year",
				"odometer",
			),
		),
		ctx.Page(),
		ctx.Sort(),
		"",
		j.GetWorkspaceId(),
	)

	ds := make([]dto.MotorGarageVehicle, len(ps))
	for i := range ps {
		ds[i] = *ps[i].ToDTO(ctx.Locale())
	}

	ctx.OK(ds, page)
}
