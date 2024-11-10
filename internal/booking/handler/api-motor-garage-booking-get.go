package handler

import (
	"errors"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/booking/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageBookingGet = metagin.Get(
	"getMotorGarageBooking",
	"Get Motor Garage Booking",
	"/motor-garage/booking/:id",
	func(ctx metagin.Context[base.RequestPathId, dto.MotorGarageBooking]) {
		booking := repo.MotorGarageBookingRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if booking == nil {
			ctx.Err(errors.New("booking not found"))
			return
		}
		ctx.OK(booking.ToDTO(ctx.Locale()))
	},
)
