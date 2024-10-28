package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/metagin/one2many"
	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice/repo"
	vehicleRepo "github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

func ApiMotorGarageInvoiceUpdate(ctx metagin.IContext[request.MotorGarageInvoiceUpdate]) {

	j := ctx.Jwt()

	if err := ctx.GetRequest().Validate(); err != "" {
		ctx.Err(err)
		return
	}

	e := repo.MotorGarageInvoiceRepo.FindByID(ctx.GetDB().Preload("Vehicle").Preload("Services").Preload("Products").Preload("Discounts"), ctx.GetRequest().ID, j.GetWorkspaceId())
	if e == nil {
		ctx.Err("invoice not found")
		return
	}

	vehicle := vehicleRepo.MotorGarageVehicleRepo.FindByID(ctx.GetDB(), ctx.GetRequest().VehicleId, j.GetWorkspaceId())
	if vehicle == nil {
		ctx.Err("vehicle not found")
		return
	}

	e = ctx.GetRequest().ToEntity(e)
	e.Vehicle = vehicle

	e = repo.MotorGarageInvoiceRepo.Save(ctx.GetDB(), e, j.GetWorkspaceId())
	if e == nil {
		ctx.InternalServerError("failed to save invoice")
		return
	}

	if len(e.Services) > 0 {
		if !one2many.HandleOne2Many(ctx.GetDB(), e.ID, "InvoiceId", e.Services, j.GetWorkspaceId()) {
			ctx.InternalServerError("failed to save invoice services")
			return
		}
	}

	if len(e.Products) > 0 {
		if !one2many.HandleOne2Many(ctx.GetDB(), e.ID, "InvoiceId", e.Products, j.GetWorkspaceId()) {
			ctx.InternalServerError("failed to save invoice products")
			return
		}
	}

	if len(e.Discounts) > 0 {
		if !one2many.HandleOne2Many(ctx.GetDB(), e.ID, "InvoiceId", e.Discounts, j.GetWorkspaceId()) {
			ctx.InternalServerError("failed to save invoice discounts")
			return
		}
	}

	ctx.OK(e.ToDTO())

}
