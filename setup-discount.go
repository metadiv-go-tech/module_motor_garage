package module_motor_garage

import (
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/discount/handler"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

func SetupDiscount() {
	metagin.Migrate(&entity.MotorGarageDiscount{})
	metagin.RegisterHandler(
		handler.ApiMotorGarageDiscountCreate,
		handler.ApiMotorGarageDiscountList,
		handler.ApiMotorGarageDiscountGet,
		handler.ApiMotorGarageDiscountUpdate,
		handler.ApiMotorGarageDiscountDelete,
	)
}
