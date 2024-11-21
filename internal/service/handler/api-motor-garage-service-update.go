package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/config"
	"github.com/metadiv-go-tech/module_motor_garage/internal/service/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

var ApiMotorGarageServiceUpdate = metagin.Put(
	"updateService",
	"Update Service",
	fmt.Sprintf("/api/%s/motor-garage/service/:id", config.SystemVersion),
	func(ctx metagin.Context[request.MotorGarageServiceUpdate, dto.MotorGarageService]) {

		s := repo.MotorGarageServiceRepo.FindById(ctx.DB(), ctx.Request().ID, ctx.WorkspaceId())
		if s == nil {
			ctx.Err(errors.New("service not found"))
			return
		}

		if errMsg := ctx.Request().Validate(); errMsg != "" {
			ctx.Err(errors.New(errMsg))
			return
		}

		s = ctx.Request().ToEntity(s)
		s = repo.MotorGarageServiceRepo.Save(ctx.DB(), s, ctx.WorkspaceId())
		if s == nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save service"))
			return
		}

		ctx.OK(s.ToDTO())
	},
)
