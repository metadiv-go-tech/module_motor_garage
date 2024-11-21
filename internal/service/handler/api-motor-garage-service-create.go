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

var ApiMotorGarageServiceCreate = metagin.Post(
	"createService",
	"Create Service",
	fmt.Sprintf("/api/%s/motor-garage/service", config.SystemVersion),
	func(ctx metagin.Context[request.MotorGarageServiceCreate, dto.MotorGarageService]) {

		s := ctx.Request().ToEntity(nil)

		if errMsg := ctx.Request().Validate(); errMsg != "" {
			ctx.Err(errors.New(errMsg))
			return
		}

		s = repo.MotorGarageServiceRepo.Save(ctx.DB(), s, ctx.WorkspaceId())
		if s == nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New("failed to save service"))
			return
		}

		ctx.OK(s.ToDTO())
	},
)
