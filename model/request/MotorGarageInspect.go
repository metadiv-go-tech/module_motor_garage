package request

import (
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

type MotorGarageInspect struct {
	ID        uint `json:"id"`
	InvoiceId uint `json:"invoice_id"`

	CustomerInstructionsAndRepairs *dto.MotorGarageCustomerInstructionsAndRepairs      `json:"customer_instructions_and_repairs"`
	RoadTest                       *dto.MotorGarageInspectRoadTest                     `json:"road_test"`
	EngineTune                     *dto.MotorGarageInspectEngineTune                   `json:"engine_tune"`
	LightChecks                    *dto.MotorGarageInspectLightChecks                  `json:"light_checks"`
	InteriorChecks                 *dto.MotorGarageInspectInteriorChecks               `json:"interior_checks"`
	UnderBody                      *dto.MotorGarageInspectUnderBody                    `json:"under_body"`
	ExhaustSystemChecks            *dto.MotorGarageInspectExhaustSystemChecks          `json:"exhaust_system_checks"`
	SuspensionSteeringSystemTest   *dto.MotorGarageInspectSuspensionSteeringSystemTest `json:"suspension_steering_system_test"`
	BreakingSystemTest             *dto.MotorGarageInspectBreakingSystemTest           `json:"breaking_system_test"`
	UnderTheBonnetTests            *dto.MotorGarageInspectUnderTheBonnetTests          `json:"under_the_bonnet_tests"`
	FinalProcedures                *dto.MotorGarageInspectFinalProcedures              `json:"final_procedures"`
}

func (r *MotorGarageInspect) ToEntity(e *entity.MotorGarageInspect) *entity.MotorGarageInspect {
	if e == nil {
		e = &entity.MotorGarageInspect{}
	}

	e.CustomerInstructionsAndRepairs = dto.InspectSectionToString(r.CustomerInstructionsAndRepairs)
	e.RoadTest = dto.InspectSectionToString(r.RoadTest)
	e.EngineTune = dto.InspectSectionToString(r.EngineTune)
	e.LightChecks = dto.InspectSectionToString(r.LightChecks)
	e.InteriorChecks = dto.InspectSectionToString(r.InteriorChecks)
	e.UnderBody = dto.InspectSectionToString(r.UnderBody)
	e.ExhaustSystemChecks = dto.InspectSectionToString(r.ExhaustSystemChecks)
	e.SuspensionSteeringSystemTest = dto.InspectSectionToString(r.SuspensionSteeringSystemTest)
	e.BreakingSystemTest = dto.InspectSectionToString(r.BreakingSystemTest)
	e.UnderTheBonnetTests = dto.InspectSectionToString(r.UnderTheBonnetTests)
	e.FinalProcedures = dto.InspectSectionToString(r.FinalProcedures)

	return e
}
