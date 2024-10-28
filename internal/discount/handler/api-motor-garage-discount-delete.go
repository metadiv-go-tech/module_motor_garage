package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/module_motor_garage/internal/discount/repo"
)

// @Summary Delete Motor Garage Discount by ID
// @Description Delete Motor Garage Discount by ID
// @Tags Motor Garage Discount
// @Param id path int true "ID"
// @Success 200
// @Failure 40001 {object} string "discount not found"
// @Router /motor-garage-discount/{id} [delete]
func ApiMotorGarageDiscountDelete(ctx metagin.IContext[base.RequestIDPath]) {
	j := ctx.Jwt()

	i := repo.MotorGarageDiscountRepo.FindByID(ctx.GetDB(), ctx.GetRequest().ID, j.GetWorkspaceId())
	if i == nil {
		ctx.Err("discount not found")
		return
	}

	if !repo.MotorGarageDiscountRepo.Delete(ctx.GetDB(), i) {
		ctx.InternalServerError("failed to delete discount")
		return
	}

	ctx.OK(nil)
}
