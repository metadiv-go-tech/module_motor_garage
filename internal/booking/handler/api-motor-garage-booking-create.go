package handler

import (
	"errors"
	"net/http"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/booking/repo"
	vehicleRepo "github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
	relationship "github.com/metadiv-go-tech/module_relationship/v2"
)

var ApiMotorGarageBookingCreate = metagin.Post(
	"createMotorGarageBooking",
	"Create Motor Garage Booking",
	"/motor-garage/booking",
	func(ctx metagin.Context[request.MotorGarageBookingCreate, dto.MotorGarageBooking]) {

		if err := ctx.Request().Validate(); err != nil {
			ctx.Err(err)
			return
		}

		booking := ctx.Request().ToEntity(nil)

		if booking.CustomerId != nil {
			customer := relationship.CustomerCaller.GetCustomerByID(ctx.DB(), *booking.CustomerId, ctx.WorkspaceId())
			if customer == nil {
				ctx.Err(errors.New("customer not found"))
				return
			}
			booking.Customer = customer
		}

		if booking.VehicleId != nil {
			vehicle := vehicleRepo.MotorGarageVehicleRepo.FindById(ctx.DB().Preload("Customer"), *booking.VehicleId, ctx.WorkspaceId())
			if vehicle == nil {
				ctx.Err(errors.New("vehicle not found"))
				return
			}
			booking.Vehicle = vehicle
			if booking.CustomerId == nil {
				booking.CustomerId = vehicle.CustomerId
			}
		}

		booking = repo.MotorGarageBookingRepo.Save(ctx.DB(), booking, ctx.WorkspaceId())
		if booking == nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save booking"))
			return
		}
		ctx.OK(booking.ToDTO(ctx.Locale()))
	},
)
