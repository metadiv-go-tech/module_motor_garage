package handler

import (
	"fmt"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metaorm/v2"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/booking/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

var ApiMotorGarageBookingList = metagin.Get(
	"listMotorGarageBooking",
	"List Motor Garage Booking",
	fmt.Sprintf("/api/%s/motor-garage/booking", config.SystemVersion),
	func(ctx metagin.Context[request.MotorGarageBookingList, []dto.MotorGarageBooking]) {

		bd := metaorm.NewAndQueryBuilder()
		if ctx.Request().From > 0 {
			bd.Add(ctx.DB().Gte("motor_garage_bookings.date_time", ctx.Request().From))
		}
		if ctx.Request().To > 0 {
			bd.Add(ctx.DB().Lte("motor_garage_bookings.date_time", ctx.Request().To))
		}
		bd.Add(ctx.DB().Or(
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
		))

		bookings, page := repo.MotorGarageBookingRepo.FindAllComplexJoined(
			ctx.DB().Joins("Customer", "Customer.ContactPerson", "Vehicle"),
			bd.Build(),
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
