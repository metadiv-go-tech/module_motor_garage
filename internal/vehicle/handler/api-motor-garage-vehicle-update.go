package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/metaorm"
	"github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
	relationship "github.com/metadiv-go-tech/module_relationship"
)

func ApiMotorGarageVehicleUpdate(ctx metagin.IContext[request.MotorGarageVehicleUpdate]) {
	j := ctx.Jwt()

	if errMsg := ctx.GetRequest().Validate(); errMsg != "" {
		ctx.Err(errMsg)
		return
	}

	v := repo.MotorGarageVehicleRepo.FindByID(ctx.GetDB(), ctx.GetRequest().ID, j.GetWorkspaceId())
	if v == nil {
		ctx.Err("vehicle not found")
		return
	}

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

	if ctx.GetRequest().Registration != "" {
		exist := repo.MotorGarageVehicleRepo.FindOne(ctx.GetDB(), metaorm.Eq("registration", ctx.GetRequest().Registration), j.GetWorkspaceId())
		if exist != nil {
			ctx.Err("registration is already in use")
			return
		}
	}

	v = ctx.GetRequest().ToEntity(v)
	v = repo.MotorGarageVehicleRepo.Save(ctx.GetDB(), v, j.GetWorkspaceId())
	if v == nil {
		ctx.InternalServerError("failed to save vehicle")
		return
	}

	ctx.OK(v.ToDTO())
}
