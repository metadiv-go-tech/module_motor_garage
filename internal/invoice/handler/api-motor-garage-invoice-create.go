package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/one2many"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	bookingRepo "github.com/metadiv-go-tech/module_motor_garage/internal/booking/repo"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
	vehicleRepo "github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

var ApiMotorGarageInvoiceCreate = metagin.Post(
	"createInvoice",
	"Create Invoice",
	fmt.Sprintf("/api/%s/motor-garage/invoice", config.SystemVersion),
	func(ctx metagin.Context[request.MotorGarageInvoiceCreate, dto.MotorGarageInvoice]) {

		if err := ctx.Request().Validate(); err != "" {
			ctx.Err(errors.New(err))
			return
		}

		vehicle := vehicleRepo.MotorGarageVehicleRepo.FindById(ctx.DB(), ctx.Request().VehicleId, ctx.WorkspaceId())
		if vehicle == nil {
			ctx.Err(errors.New("vehicle not found"))
			return
		}

		e := ctx.Request().ToEntity(nil)
		e.Vehicle = vehicle

		if e.BookingId != nil {
			booking := bookingRepo.MotorGarageBookingRepo.FindById(ctx.DB().Preload("Invoice"), *e.BookingId, ctx.WorkspaceId())
			if booking == nil {
				ctx.Err(errors.New("booking not found"))
				return
			}
			e.Booking = booking
		}

		if ctx.Request().Inspect != nil {
			inspect := ctx.Request().Inspect.ToEntity(nil)
			e.Inspect = inspect
		}

		e = repo.MotorGarageInvoiceRepo.Save(ctx.DB(), e, ctx.WorkspaceId())
		if e == nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save invoice"))
			return
		}

		if e.Inspect != nil {
			inspect := e.Inspect
			inspect.InvoiceId = e.ID
			inspect = repo.InspectRepo.Save(ctx.DB(), inspect, ctx.WorkspaceId())
		}

		if !one2many.HandleOne2Many(ctx.DB(), e.ID, "InvoiceId", e.Services, ctx.WorkspaceId()) {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save invoice services"))
			return
		}

		if !one2many.HandleOne2Many(ctx.DB(), e.ID, "InvoiceId", e.Products, ctx.WorkspaceId()) {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save invoice products"))
			return
		}

		if !one2many.HandleOne2Many(ctx.DB(), e.ID, "InvoiceId", e.Discounts, ctx.WorkspaceId()) {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save invoice discounts"))
			return
		}

		ctx.OK(e.ToDTO(ctx.Locale()))
	},
)
