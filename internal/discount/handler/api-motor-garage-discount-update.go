package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/module_motor_garage/internal/discount/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/request"
)

// @Summary Update Motor Garage Discount by ID
// @Description Update Motor Garage Discount
// @Tags Motor Garage Discount
// @Accept  json
// @Param request body request.MotorGarageDiscountUpdate true "request"
// @Success 200 {object} dto.MotorGarageDiscount
// @Failure 40001 {object} string "discount not found"
// @Router /motor-garage-discount [put]
func ApiMotorGarageDiscountUpdate(ctx metagin.IContext[request.MotorGarageDiscountUpdate]) {

	j := ctx.Jwt()

	d := repo.MotorGarageDiscountRepo.FindByID(ctx.GetDB(), ctx.GetRequest().ID, j.GetWorkspaceId())
	if d == nil {
		ctx.Err("discount not found")
		return
	}

	d = ctx.GetRequest().ToEntity(d)
	d = repo.MotorGarageDiscountRepo.Save(ctx.GetDB(), d, j.GetWorkspaceId())
	if d == nil {
		ctx.InternalServerError("failed to save discount")
		return
	}

	ctx.OK(d.ToDTO())

}
