package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
	"github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/service"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
	relationship "github.com/metadiv-go-tech/module_relationship"
)

func ApiMotorGarageVehicleCreate(ctx metagin.IContext[request.MotorGarageVehicleCreate]) {

	j := ctx.Jwt()

	if errMsg := ctx.GetRequest().Validate(); errMsg != "" {
		ctx.Err(errMsg)
		return
	}

	v := ctx.GetRequest().ToEntity(nil)

	if ctx.GetRequest().CustomerId != nil && *ctx.GetRequest().CustomerId > 0 {
		customer := relationship.CustomerCaller.GetCustomerByID(
			ctx.GetDB(),
			*ctx.GetRequest().CustomerId,
			j.GetWorkspaceId(),
		)
		if customer == nil {
			ctx.Err("customer not found")
			return
		}
		v.CustomerId = ctx.GetRequest().CustomerId
		v.Customer = customer
	}

	if service.VehicleService.CheckRegistration(ctx.GetDB(), ctx.GetRequest().Registration, j.GetWorkspaceId(), 0) {
		ctx.Err("registration is already in use")
		return
	}

	v = repo.MotorGarageVehicleRepo.Save(ctx.GetDB(), v, j.GetWorkspaceId())
	if v == nil {
		ctx.InternalServerError("failed to save vehicle")
		return
	}

	ctx.OK(v)
}
