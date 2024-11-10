package handler

import (
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/booking/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageBookingList = metagin.Get(
	"listMotorGarageBooking",
	"List Motor Garage Booking",
	"/motor-garage/booking",
	func(ctx metagin.Context[base.RequestListing, []dto.MotorGarageBooking]) {
		bookings, page := repo.MotorGarageBookingRepo.FindAllComplexJoined(
			ctx.DB().Joins("Customer", "Customer.ContactPerson", "Vehicle"),
			ctx.DB().Or(
				ctx.Request().BuildDecryptedSimilarClause(
					"motor_garage_bookings.note",
					"motor_garage_bookings.requirement",
					"Customer.display_name",
					"Customer.company_name",
					"Customer__ContactPerson.first_name",
					"Customer__ContactPerson.last_name",
					"Customer__ContactPerson.email",
					"Customer__ContactPerson.phone",
					"Vehicle.rego",
					"Vehicle.vin",
					"Vehicle.registration",
				),
				ctx.Request().BuildSimilarClause(
					"Vehicle.name",
				),
			),
			"motor_garage_bookings",
			ctx.Page(),
			ctx.Sort(),
			ctx.WorkspaceId(),
		)

		dtos := make([]dto.MotorGarageBooking, len(bookings))
		for i, booking := range bookings {
			dtos[i] = *booking.ToDTO(ctx.Locale())
		}
		ctx.OK(&dtos, page)
	},
)
