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

var ApiMotorGarageVehicleCreate = metagin.Post(
	"createVehicle",
	"Create Vehicle",
	"/motor-garage/vehicle",
	func(ctx metagin.Context[request.MotorGarageVehicleCreate, dto.MotorGarageVehicle]) {
		j := ctx.Jwt()

		if errMsg := ctx.Request().Validate(); errMsg != "" {
			ctx.Err(errors.New(errMsg))
			return
		}

		v := ctx.Request().ToEntity(nil)

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

		if service.VehicleService.CheckRegistration(ctx.DB(), ctx.Request().Registration, j.GetWorkspaceId(), 0) {
			ctx.Err(errors.New("registration is already in use"))
			return
		}

		v = repo.MotorGarageVehicleRepo.Save(ctx.DB(), v, j.GetWorkspaceId())
		if v == nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save vehicle"))
			return
		}

		ctx.OK(v.ToDTO(ctx.Locale()))
	},
)
