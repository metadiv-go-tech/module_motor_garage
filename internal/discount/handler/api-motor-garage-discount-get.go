package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/discount/repo"
)

// @Summary Get Motor Garage Discount by ID
// @Description Get Motor Garage Discount by ID
// @Tags Motor Garage Discount
// @Param id path int true "ID"
// @Success 200 {object} dto.MotorGarageDiscount
// @Failure 40001 {object} string "discount not found"
// @Router /motor-garage-discount/{id} [get]
func ApiMotorGarageDiscountGet(ctx metagin.IContext[base.RequestIDPath]) {

	j := ctx.Jwt()

	d := repo.MotorGarageDiscountRepo.FindByID(ctx.GetDB(), ctx.GetRequest().ID, j.GetWorkspaceId())
	if d == nil {
		ctx.Err("discount not found")
		return
	}

	ctx.OK(d.ToDTO())

}
