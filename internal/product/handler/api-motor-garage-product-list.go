package handler

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/metaorm"
	"github.com/metadiv-go-tech/module_motor_garage/internal/product/repo"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

func ApiMotorGarageProductList(ctx metagin.IContext[base.RequestListing]) {
	j := ctx.Jwt()
	ps, page := repo.MotorGarageProductRepo.FindAllComplex(
		ctx.GetDB(),
		metaorm.Or(
			ctx.GetRequest().BuildDecryptedSimilarClause(
				"name",
				"description",
			),
			ctx.GetRequest().BuildSimilarClause(
				"price",
				"price_after_tax",
			),
		),
		ctx.Page(),
		ctx.Sort(),
		"",
		j.GetWorkspaceId(),
	)
	ds := make([]dto.MotorGarageProduct, len(ps))
	for i := range ps {
		ds[i] = *ps[i].ToDTO()
	}
	ctx.OK(ds, page)
}
