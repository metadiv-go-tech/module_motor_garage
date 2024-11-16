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

	RoadTest                     string `json:"road_test"`
	EngineTune                   string `json:"engine_tune"`
	LightChecks                  string `json:"light_checks"`
	InteriorChecks               string `json:"interior_checks"`
	UnderBody                    string `json:"under_body"`
	ExhaustSystemChecks          string `json:"exhaust_system_checks"`
	SuspensionSteeringSystemTest string `json:"suspension_steering_system_test"`
	BreakingSystemTest           string `json:"breaking_system_test"`
	UnderTheBonnetTests          string `json:"under_the_bonnet_tests"`
	FinalProcedures              string `json:"final_procedures"`
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
	if e.UnderBody != "" {
		d.UnderBody = &dto.MotorGarageInspectUnderBody{}
		json.Unmarshal([]byte(e.UnderBody), d.UnderBody)
	}
	if e.ExhaustSystemChecks != "" {
		d.ExhaustSystemChecks = &dto.MotorGarageInspectExhaustSystemChecks{}
		json.Unmarshal([]byte(e.ExhaustSystemChecks), d.ExhaustSystemChecks)
	}
	if e.SuspensionSteeringSystemTest != "" {
		d.SuspensionSteeringSystemTest = &dto.MotorGarageInspectSuspensionSteeringSystemTest{}
		json.Unmarshal([]byte(e.SuspensionSteeringSystemTest), d.SuspensionSteeringSystemTest)
	}
	if e.BreakingSystemTest != "" {
		d.BreakingSystemTest = &dto.MotorGarageInspectBreakingSystemTest{}
		json.Unmarshal([]byte(e.BreakingSystemTest), d.BreakingSystemTest)
	}
	if e.UnderTheBonnetTests != "" {
		d.UnderTheBonnetTests = &dto.MotorGarageInspectUnderTheBonnetTests{}
		json.Unmarshal([]byte(e.UnderTheBonnetTests), d.UnderTheBonnetTests)
	}
	if e.FinalProcedures != "" {
		d.FinalProcedures = &dto.MotorGarageInspectFinalProcedures{}
		json.Unmarshal([]byte(e.FinalProcedures), d.FinalProcedures)
	}

	return d
}
