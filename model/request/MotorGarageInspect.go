package request

import (
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

type MotorGarageInspect struct {
	ID        uint `json:"id"`
	InvoiceId uint `json:"invoice_id"`

	RoadTest *dto.MotorGarageInspectRoadTest `json:"road_test"`
}

func (r *MotorGarageInspect) ToEntity(e *entity.MotorGarageInspect) *entity.MotorGarageInspect {
	if e == nil {
		e = &entity.MotorGarageInspect{}
	}
	e.RoadTest = r.RoadTest.ToString()
	return e
}
