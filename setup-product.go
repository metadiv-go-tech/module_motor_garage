package module_motor_garage

import (
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/product/handler"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

func SetupProduct() {
	metagin.Migrate(&entity.MotorGarageProduct{})
	metagin.RegisterHandler(
		handler.ApiMotorGarageProductCreate,
		handler.ApiMotorGarageProductList,
		handler.ApiMotorGarageProductGet,
		handler.ApiMotorGarageProductUpdate,
		handler.ApiMotorGarageProductDelete,
	)
}
