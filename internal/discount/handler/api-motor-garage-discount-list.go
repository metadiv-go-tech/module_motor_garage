package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/metaorm"
	"github.com/metadiv-go-tech/module_motor_garage/internal/discount/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

// @Summary List Motor Garage Discounts
// @Description List Motor Garage Discounts
// @Tags Motor Garage Discount
// @Param name query string false "Name"
// @Param description query string false "Description"
// @Param page query int false "page"
// @Param size query int false "size"
// @Param field query string false "field"
// @Param asc query bool false "sort asc"
// @Param keyword query string false "keyword"
// @Success 200 {array} dto.MotorGarageDiscount
// @Router /motor-garage-discount [get]
func ApiMotorGarageDiscountList(ctx metagin.IContext[base.RequestListing]) {

	j := ctx.Jwt()

	ps, page := repo.MotorGarageDiscountRepo.FindAllComplex(
		ctx.GetDB(),
		metaorm.Or(
			ctx.GetRequest().BuildSimilarClause(
				"name",
				"description",
			),
		),
		ctx.Page(),
		ctx.Sort(),
		"",
		j.GetWorkspaceId(),
	)

	ds := make([]dto.MotorGarageDiscount, len(ps))
	for i := range ps {
		ds[i] = *ps[i].ToDTO()
	}

	ctx.OK(ds, page)
}
