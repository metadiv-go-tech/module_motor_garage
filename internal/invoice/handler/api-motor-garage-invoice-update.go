package handler

import (
	"errors"
	"net/http"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/one2many"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
	vehicleRepo "github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

var ApiMotorGarageInvoiceUpdate = metagin.Put(
	"updateInvoice",
	"Update Invoice",
	"/motor-garage/invoice/:id",
	func(ctx metagin.Context[request.MotorGarageInvoiceUpdate, dto.MotorGarageInvoice]) {

		j := ctx.Jwt()

		if err := ctx.Request().Validate(); err != "" {
			ctx.Err(errors.New(err))
			return
		}

		e := repo.MotorGarageInvoiceRepo.FindById(ctx.DB(), ctx.Request().ID, j.GetWorkspaceId())
		if e == nil {
			ctx.Err(errors.New("invoice not found"))
			return
		}

		vehicle := vehicleRepo.MotorGarageVehicleRepo.FindById(ctx.DB(), ctx.Request().VehicleId, j.GetWorkspaceId())
		if vehicle == nil {
			ctx.Err(errors.New("vehicle not found"))
			return
		}

		e = ctx.Request().ToEntity(e)
		e.Vehicle = vehicle

		if !one2many.HandleOne2Many(ctx.DB(), e.ID, "InvoiceId", e.Services, j.GetWorkspaceId()) {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save invoice services"))
			return
		}

		if !one2many.HandleOne2Many(ctx.DB(), e.ID, "InvoiceId", e.Products, j.GetWorkspaceId()) {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save invoice products"))
			return
		}

		if !one2many.HandleOne2Many(ctx.DB(), e.ID, "InvoiceId", e.Discounts, j.GetWorkspaceId()) {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save invoice discounts"))
			return
		}

		e.Services = nil
		e.Products = nil
		e.Discounts = nil
		e = repo.MotorGarageInvoiceRepo.Save(ctx.DB(), e, j.GetWorkspaceId())
		if e == nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save invoice"))
			return
		}

		e = repo.MotorGarageInvoiceRepo.FindById(ctx.DB().Preload("Products", "Services", "Discounts"), ctx.Request().ID, j.GetWorkspaceId())
		ctx.OK(e.ToDTO(ctx.Locale()))

	},
)
