package handler

import (
	"errors"
	"net/http"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
	"github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/service"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
	relationship "github.com/metadiv-go-tech/module_relationship/v2"
)

var ApiMotorGarageVehicleUpdate = metagin.Put(
	"updateVehicle",
	"Update Vehicle",
	"/motor-garage/vehicle/:id",
	func(ctx metagin.Context[request.MotorGarageVehicleUpdate, dto.MotorGarageVehicle]) {
		j := ctx.Jwt()

		if errMsg := ctx.Request().Validate(); errMsg != "" {
			ctx.Err(errors.New(errMsg))
			return
		}

		v := repo.MotorGarageVehicleRepo.FindById(ctx.DB(), ctx.Request().ID, j.GetWorkspaceId())
		if v == nil {
			ctx.Err(errors.New("vehicle not found"))
			return
		}

		if ctx.Request().CustomerId != nil && *ctx.Request().CustomerId > 0 {
			customer := relationship.CustomerCaller.GetCustomerByID(
				ctx.DB(),
				*ctx.Request().CustomerId,
				j.GetWorkspaceId(),
			)
			if customer == nil {
				ctx.Err(errors.New("customer not found"))
				return
			}
			v.CustomerId = ctx.Request().CustomerId
			v.Customer = customer
		}

		if ctx.Request().Registration != "" {
			if service.VehicleService.CheckRegistration(ctx.DB(), ctx.Request().Registration, j.GetWorkspaceId(), ctx.Request().ID) {
				ctx.Err(errors.New("registration is already in use"))
				return
			}
		}

		v = ctx.Request().ToEntity(v)
		v = repo.MotorGarageVehicleRepo.Save(ctx.DB(), v, j.GetWorkspaceId())
		if v == nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save vehicle"))
			return
		}

		ctx.OK(v.ToDTO(ctx.Locale()))
	},
)
