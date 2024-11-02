package handler

import (
	"errors"
	"net/http"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/service/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

var ApiMotorGarageServiceCreate = metagin.Post(
	"createService",
	"Create Service",
	"/motor-garage/service",
	func(ctx metagin.Context[request.MotorGarageServiceCreate, dto.MotorGarageService]) {
		j := ctx.Jwt()

		s := ctx.Request().ToEntity(nil)

		if errMsg := ctx.Request().Validate(); errMsg != "" {
			ctx.Err(errors.New(errMsg))
			return
		}

		s = repo.MotorGarageServiceRepo.Save(ctx.DB(), s, j.GetWorkspaceId())
		if s == nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save service"))
			return
		}

		ctx.OK(s.ToDTO())
	},
)
