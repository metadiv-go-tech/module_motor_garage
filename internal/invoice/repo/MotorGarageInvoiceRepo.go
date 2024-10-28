package repo

import (
	"github.com/metadiv-go-tech/metagin/base"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

var MotorGarageInvoiceRepo = new(motorGarageInvoiceRepo)

type motorGarageInvoiceRepo struct {
	base.Repository[entity.MotorGarageInvoice]
}
