package handler

import (
	"errors"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/booking/repo"
	invoiceRepo "github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
	vehicleRepo "github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
	relationship "github.com/metadiv-go-tech/module_relationship/v2"
)

var ApiMotorGarageBookingUpdate = metagin.Put(
	"updateMotorGarageBooking",
	"Update Motor Garage Booking",
	"/motor-garage/booking/:id",
	func(ctx metagin.Context[request.MotorGarageBookingUpdate, dto.MotorGarageBooking]) {

		if err := ctx.Request().Validate(); err != nil {
			ctx.Err(err)
			return
		}

		booking := repo.MotorGarageBookingRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if booking == nil {
			ctx.Err(errors.New("booking not found"))
			return
		}

		ctx.Request().ToEntity(booking)

		if booking.CustomerId != nil {
			customer := relationship.CustomerCaller.GetCustomerByID(ctx.DB(), *booking.CustomerId, ctx.WorkspaceId())
			if customer == nil {
				ctx.Err(errors.New("customer not found"))
				return
			}
			booking.Customer = customer
		}

		if booking.VehicleId != nil {
			vehicle := vehicleRepo.MotorGarageVehicleRepo.FindById(ctx.DB(), *booking.VehicleId, ctx.WorkspaceId())
			if vehicle == nil {
				ctx.Err(errors.New("vehicle not found"))
				return
			}
			booking.Vehicle = vehicle
			if booking.CustomerId == nil {
				booking.CustomerId = vehicle.CustomerId
			}
		}

		if booking.InvoiceId != nil {
			invoice := invoiceRepo.MotorGarageInvoiceRepo.FindById(ctx.DB(), *booking.InvoiceId, ctx.WorkspaceId())
			if invoice == nil {
				ctx.Err(errors.New("invoice not found"))
				return
			}
			booking.Invoice = invoice
		}

		booking = repo.MotorGarageBookingRepo.Save(ctx.DB(), booking, ctx.WorkspaceId())
		if booking == nil {
			ctx.Err(errors.New("failed to update booking"))
			return
		}
		ctx.OK(booking.ToDTO(ctx.Locale()))
	},
)
