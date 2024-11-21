package handler

import (
	"errors"
	"fmt"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/booking/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

var ApiMotorGarageBookingDelete = metagin.Delete(
	"deleteMotorGarageBooking",
	"Delete Motor Garage Booking",
	fmt.Sprintf("/api/%s/motor-garage/booking/:id", config.SystemVersion),
	func(ctx metagin.Context[base.RequestPathId, dto.MotorGarageBooking]) {

		booking := repo.MotorGarageBookingRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if booking == nil {
			ctx.Err(errors.New("booking not found"))
			return
		}
		repo.MotorGarageBookingRepo.Delete(ctx.DB(), booking)
		ctx.OK(nil)
	},
)
