package handler

import (
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/inspect/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageInspectList = metagin.Get(
	"listInspect",
	"List motor inspect",
	"/motor-garage/inspect",
	func(ctx metagin.Context[base.RequestListing, []dto.MotorGarageInspect]) {
		inspects, page := repo.InspectRepo.FindAllComplexJoined(
			ctx.DB().Joins(
				"Invoice", "Invoice.Vehicle",
				"Invoice.Vehicle.Customer",
				"Invoice.Vehicle.Customer.ContactPerson"),
			ctx.DB().Or(
				ctx.Request().BuildDecryptedSimilarClause(
					"Vehicle.rego",
					"Vehicle.vin",
					"Vehicle.registration",
					"Customer.company_name",
					"Customer.display_name",
					"Customer.ContactPerson.first_name",
					"Customer.ContactPerson.last_name",
					"Customer.ContactPerson.email",
					"Customer.ContactPerson.phone",
				),
				ctx.Request().BuildSimilarClause(
					"Vehicle.name",
					"Customer.code",
				),
			),
			"motor_garage_inspects",
			ctx.Page(),
			ctx.Sort(),
			ctx.WorkspaceId(),
		)

		dtos := make([]dto.MotorGarageInspect, len(inspects))
		for i, inspect := range inspects {
			dtos[i] = *inspect.ToDTO(ctx.Locale())
		}
		ctx.OK(&dtos, page)
	},
)
