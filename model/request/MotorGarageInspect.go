package request

import (
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

type MotorGarageInspect struct {
	ID        uint `json:"id"`
	InvoiceId uint `json:"invoice_id"`

	RoadTest       *dto.MotorGarageInspectRoadTest       `json:"road_test"`
	EngineTune     *dto.MotorGarageInspectEngineTune     `json:"engine_tune"`
	LightChecks    *dto.MotorGarageInspectLightChecks    `json:"light_checks"`
	InteriorChecks *dto.MotorGarageInspectInteriorChecks `json:"interior_checks"`
}

func (r *MotorGarageInspect) ToEntity(e *entity.MotorGarageInspect) *entity.MotorGarageInspect {
	if e == nil {
		e = &entity.MotorGarageInspect{}
	}
	e.RoadTest = dto.InspectSectionToString(r.RoadTest)
	e.EngineTune = dto.InspectSectionToString(r.EngineTune)
	e.LightChecks = dto.InspectSectionToString(r.LightChecks)
	e.InteriorChecks = dto.InspectSectionToString(r.InteriorChecks)
	return e
}
