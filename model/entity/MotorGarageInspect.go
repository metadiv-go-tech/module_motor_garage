package entity

import (
	"encoding/json"

	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

type MotorGarageInspect struct {
	base.Model
	base.ModelWorkspace

	InvoiceId uint                `json:"invoice_id"`
	Invoice   *MotorGarageInvoice `json:"invoice" gorm:"foreignKey:InvoiceId"`

	RoadTest       string `json:"road_test"`
	EngineTune     string `json:"engine_tune"`
	LightChecks    string `json:"light_checks"`
	InteriorChecks string `json:"interior_checks"`
}

func (e *MotorGarageInspect) ToDTO(locale string) *dto.MotorGarageInspect {
	d := &dto.MotorGarageInspect{
		ID:        e.ID,
		InvoiceId: e.InvoiceId,
	}
	if e.Invoice != nil {
		d.Invoice = e.Invoice.ToDTO(locale)
	}
	if e.RoadTest != "" {
		d.RoadTest = &dto.MotorGarageInspectRoadTest{}
		json.Unmarshal([]byte(e.RoadTest), d.RoadTest)
	}
	if e.EngineTune != "" {
		d.EngineTune = &dto.MotorGarageInspectEngineTune{}
		json.Unmarshal([]byte(e.EngineTune), d.EngineTune)
	}
	if e.LightChecks != "" {
		d.LightChecks = &dto.MotorGarageInspectLightChecks{}
		json.Unmarshal([]byte(e.LightChecks), d.LightChecks)
	}
	if e.InteriorChecks != "" {
		d.InteriorChecks = &dto.MotorGarageInspectInteriorChecks{}
		json.Unmarshal([]byte(e.InteriorChecks), d.InteriorChecks)
	}

	return d
}
