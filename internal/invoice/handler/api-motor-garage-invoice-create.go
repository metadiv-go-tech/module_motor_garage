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
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/service"
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

		if ctx.Request().BookingId != nil {
			booking := bookingRepo.MotorGarageBookingRepo.FindById(ctx.DB().Preload("Invoice"), *ctx.Request().BookingId, ctx.WorkspaceId())
			if booking == nil {
				ctx.Err(errors.New("booking not found"))
				return
			}
			booking.InvoiceId = &e.ID
			booking = bookingRepo.MotorGarageBookingRepo.Save(ctx.DB(), booking, ctx.WorkspaceId())
			e.Booking = booking
		}

		services := e.Services
		products := e.Products
		discounts := e.Discounts
		e.Services = nil
		e.Products = nil
		e.Discounts = nil

		e = repo.MotorGarageInvoiceRepo.Save(ctx.DB(), e, ctx.WorkspaceId())
		if e == nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save invoice"))
			return
		}

		if ctx.Request().Inspect != nil {
			inspect := ctx.Request().Inspect.ToEntity(nil)
			inspect.InvoiceId = e.ID
			inspect = repo.InspectRepo.Save(ctx.DB(), inspect, ctx.WorkspaceId())
			e.Inspect = inspect
		}

		if !one2many.HandleOne2Many(ctx.DB(), e.ID, "InvoiceId", services, ctx.WorkspaceId()) {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save invoice services"))
			return
		}

		if !one2many.HandleOne2Many(ctx.DB(), e.ID, "InvoiceId", products, ctx.WorkspaceId()) {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save invoice products"))
			return
		}

		if !one2many.HandleOne2Many(ctx.DB(), e.ID, "InvoiceId", discounts, ctx.WorkspaceId()) {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save invoice discounts"))
			return
		}

		e = repo.MotorGarageInvoiceRepo.FindById(ctx.DB().Preload("Products", "Services", "Discounts"), e.ID, ctx.WorkspaceId())
		if e == nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to load invoice"))
			return
		}

		e.Total = service.InvoiceService.CalculateInvoiceTotal(e)
		e = repo.MotorGarageInvoiceRepo.Save(ctx.DB(), e, ctx.WorkspaceId())
		if e == nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save invoice total"))
			return
		}

		ctx.OK(e.ToDTO(ctx.Locale()))
	},
)
