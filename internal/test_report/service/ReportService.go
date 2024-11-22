package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/metadiv-go-tech/module_motor_garage/internal/test_report/template"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

var ReportService = new(reportService)

type reportService struct{}

func (s *reportService) GenerateReport(invoice *entity.MotorGarageInvoice, locale string) string {

	html := template.TestReportTemplate

	html = strings.Replace(html, "{{invoice_number}}", fmt.Sprintf("#%d", invoice.ID), -1)
	html = strings.Replace(html, "{{invoice_date}}", time.Unix(invoice.Date, 0).Format("02/01/2006"), -1)

	if invoice.Vehicle != nil && invoice.Vehicle.Customer != nil && invoice.Vehicle.Customer.ContactPerson != nil {
		contactPerson := invoice.Vehicle.Customer.ContactPerson.ToDTO()
		html = strings.Replace(html, "{{customer_name}}", contactPerson.FirstName+" "+contactPerson.LastName, -1)
	} else {
		html = strings.Replace(html, "{{customer_name}}", "-", -1)
	}

	if invoice.Vehicle != nil {
		html = strings.Replace(html, "{{vehicle_name}}", invoice.Vehicle.Name, -1)
		html = strings.Replace(html, "{{vehicle_year}}", fmt.Sprintf("%d", invoice.Vehicle.Year), -1)
		html = strings.Replace(html, "{{vehicle_odometer}}", fmt.Sprintf("%d", invoice.Vehicle.Odometer), -1)
		vehicleDto := invoice.Vehicle.ToDTO(locale)
		html = strings.Replace(html, "{{vehicle_registration}}", vehicleDto.Registration, -1)
	} else {
		html = strings.Replace(html, "{{vehicle_name}}", "-", -1)
		html = strings.Replace(html, "{{vehicle_year}}", "-", -1)
		html = strings.Replace(html, "{{vehicle_odometer}}", "-", -1)
	}

	if invoice.Inspect != nil {
		inspectDto := invoice.Inspect.ToDTO(locale)

		// Customer Instructions and Repairs
		if inspectDto.CustomerInstructionsAndRepairs != nil {
			html = strings.Replace(html, "{{tune}}", dto.GetCheckboxCheckStatusString(inspectDto.CustomerInstructionsAndRepairs.Instruction1), -1)
			html = strings.Replace(html, "{{injector}}", dto.GetCheckboxCheckStatusString(inspectDto.CustomerInstructionsAndRepairs.Instruction2), -1)
			html = strings.Replace(html, "{{diesel_service}}", dto.GetCheckboxCheckStatusString(inspectDto.CustomerInstructionsAndRepairs.Instruction3), -1)
			html = strings.Replace(html, "{{standard}}", dto.GetCheckboxCheckStatusString(inspectDto.CustomerInstructionsAndRepairs.Instruction4), -1)
			html = strings.Replace(html, "{{timing_belt}}", dto.GetCheckboxCheckStatusString(inspectDto.CustomerInstructionsAndRepairs.Instruction5), -1)
			html = strings.Replace(html, "{{diagnose}}", dto.GetCheckboxCheckStatusString(inspectDto.CustomerInstructionsAndRepairs.Instruction6), -1)
			html = strings.Replace(html, "{{major_service}}", dto.GetCheckboxCheckStatusString(inspectDto.CustomerInstructionsAndRepairs.Instruction7), -1)
			html = strings.Replace(html, "{{exhaust}}", dto.GetCheckboxCheckStatusString(inspectDto.CustomerInstructionsAndRepairs.Instruction8), -1)
			html = strings.Replace(html, "{{auto_service}}", dto.GetCheckboxCheckStatusString(inspectDto.CustomerInstructionsAndRepairs.Instruction9), -1)
			html = strings.Replace(html, "{{minor_service}}", dto.GetCheckboxCheckStatusString(inspectDto.CustomerInstructionsAndRepairs.Instruction10), -1)
			html = strings.Replace(html, "{{suspension}}", dto.GetCheckboxCheckStatusString(inspectDto.CustomerInstructionsAndRepairs.Instruction11), -1)
			html = strings.Replace(html, "{{vehicle_check}}", dto.GetCheckboxCheckStatusString(inspectDto.CustomerInstructionsAndRepairs.Instruction12), -1)
			html = strings.Replace(html, "{{wiper_disc_sce}}", dto.GetCheckboxCheckStatusString(inspectDto.CustomerInstructionsAndRepairs.Instruction13), -1)
			html = strings.Replace(html, "{{logbook_service}}", dto.GetCheckboxCheckStatusString(inspectDto.CustomerInstructionsAndRepairs.Instruction14), -1)
			html = strings.Replace(html, "{{cab_over_schg}}", dto.GetCheckboxCheckStatusString(inspectDto.CustomerInstructionsAndRepairs.Instruction15), -1)
			html = strings.Replace(html, "{{authority_to_proceed}}", dto.GetCheckboxCheckStatusString(inspectDto.CustomerInstructionsAndRepairs.Instruction16), -1)
			html = strings.Replace(html, "{{prime_item_of_concern}}", inspectDto.CustomerInstructionsAndRepairs.PrimeItemOfConcern, -1)
		}

		// 1. Road Test
		if inspectDto.RoadTest != nil {
			html = strings.Replace(html, "{{fit_seat_cover_and_floor_mat}}", inspectDto.RoadTest.Item1.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{check_oil_water_and_tyre_condition}}", inspectDto.RoadTest.Item2.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{obvious_vehicle_damage}}", inspectDto.RoadTest.Item3.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{road_test_max_speed_reached}}", inspectDto.RoadTest.Item4.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{road_test_max_speed}}", inspectDto.RoadTest.Item4.Speed, -1)
			html = strings.Replace(html, "{{air_conditioning_climate_heater_controls}}", inspectDto.RoadTest.Item5.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{air_conditioning_temperature_check}}", inspectDto.RoadTest.Item6.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{handbrake_operation}}", inspectDto.RoadTest.Item7.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{footbrake_operation}}", inspectDto.RoadTest.Item8.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{clutch_gearbox_operation}}", inspectDto.RoadTest.Item9.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{automatic_transmission_operation}}", inspectDto.RoadTest.Item10.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{starter_inhibitor_switch}}", inspectDto.RoadTest.Item11.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{requires_injector_service}}", inspectDto.RoadTest.Item12.GetPassOrFailString(), -1)
		}

		// 2. Engine Tune
		if inspectDto.EngineTune != nil {
			html = strings.Replace(html, "{{battery_load_test_before}}", inspectDto.EngineTune.Item13.Before, -1)
			html = strings.Replace(html, "{{battery_load_test_after}}", inspectDto.EngineTune.Item13.After, -1)
			html = strings.Replace(html, "{{cranking_voltage_before}}", inspectDto.EngineTune.Item14.Before, -1)
			html = strings.Replace(html, "{{cranking_voltage_after}}", inspectDto.EngineTune.Item14.After, -1)
			html = strings.Replace(html, "{{charging_voltage_before}}", inspectDto.EngineTune.Item15.Before, -1)
			html = strings.Replace(html, "{{charging_voltage_after}}", inspectDto.EngineTune.Item15.After, -1)
			html = strings.Replace(html, "{{ignition_timing_before}}", inspectDto.EngineTune.Item16.Before, -1)
			html = strings.Replace(html, "{{ignition_timing_after}}", inspectDto.EngineTune.Item16.After, -1)
			html = strings.Replace(html, "{{coil_condenser_before}}", inspectDto.EngineTune.Item17.Before, -1)
			html = strings.Replace(html, "{{coil_condenser_after}}", inspectDto.EngineTune.Item17.After, -1)
			html = strings.Replace(html, "{{idle_speed_before}}", inspectDto.EngineTune.Item18.Before, -1)
			html = strings.Replace(html, "{{idle_speed_after}}", inspectDto.EngineTune.Item18.After, -1)
			html = strings.Replace(html, "{{injectors_secure_before}}", inspectDto.EngineTune.Item19.Before, -1)
			html = strings.Replace(html, "{{injectors_secure_after}}", inspectDto.EngineTune.Item19.After, -1)
			html = strings.Replace(html, "{{fuel_lines_before}}", inspectDto.EngineTune.Item20.Before, -1)
			html = strings.Replace(html, "{{fuel_lines_after}}", inspectDto.EngineTune.Item20.After, -1)
			html = strings.Replace(html, "{{ht_leads_before}}", inspectDto.EngineTune.Item21.Before, -1)
			html = strings.Replace(html, "{{ht_leads_after}}", inspectDto.EngineTune.Item21.After, -1)
			html = strings.Replace(html, "{{egr_valve_before}}", inspectDto.EngineTune.Item22.Before, -1)
			html = strings.Replace(html, "{{egr_valve_after}}", inspectDto.EngineTune.Item22.After, -1)
			html = strings.Replace(html, "{{distributor_cap_before}}", inspectDto.EngineTune.Item23.Before, -1)
			html = strings.Replace(html, "{{distributor_cap_after}}", inspectDto.EngineTune.Item23.After, -1)
			html = strings.Replace(html, "{{rotor_before}}", inspectDto.EngineTune.Item24.Before, -1)
			html = strings.Replace(html, "{{rotor_after}}", inspectDto.EngineTune.Item24.After, -1)
			html = strings.Replace(html, "{{pcv_system_before}}", inspectDto.EngineTune.Item25.Before, -1)
			html = strings.Replace(html, "{{pcv_system_after}}", inspectDto.EngineTune.Item25.After, -1)
			html = strings.Replace(html, "{{battery_cables_levels}}", dto.GetCheckboxCheckStatusString(inspectDto.EngineTune.Item26.Levels), -1)
			html = strings.Replace(html, "{{battery_cables_before}}", inspectDto.EngineTune.Item26.Before, -1)
			html = strings.Replace(html, "{{battery_cables_after}}", inspectDto.EngineTune.Item26.After, -1)
			html = strings.Replace(html, "{{plugs_replaced_before}}", inspectDto.EngineTune.Item27.Before, -1)
			html = strings.Replace(html, "{{plugs_replaced_after}}", inspectDto.EngineTune.Item27.After, -1)

			for i := 0; i < 8; i++ {
				placeholder := fmt.Sprintf("{{compression_%v}}", i+1)

				if i < len(inspectDto.EngineTune.Item28) {
					html = strings.Replace(html, placeholder, inspectDto.EngineTune.Item28[i], -1)
				} else {
					html = strings.Replace(html, placeholder, "", -1)

				}
			}

		}

		// 3. Light checks
		if inspectDto.LightChecks != nil {
			html = strings.Replace(html, "{{headlight_operation_hi_l}}", dto.GetCheckboxCheckStatusString(inspectDto.LightChecks.Item29.Hi_L), -1)
			html = strings.Replace(html, "{{headlight_operation_hi_r}}", dto.GetCheckboxCheckStatusString(inspectDto.LightChecks.Item29.Hi_R), -1)
			html = strings.Replace(html, "{{headlight_operation_Lo_l}}", dto.GetCheckboxCheckStatusString(inspectDto.LightChecks.Item29.Lo_L), -1)
			html = strings.Replace(html, "{{headlight_operation_Lo_r}}", dto.GetCheckboxCheckStatusString(inspectDto.LightChecks.Item29.Lo_R), -1)
			html = strings.Replace(html, "{{headlight_operation}}", inspectDto.LightChecks.Item29.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{high_beam_indicator}}", inspectDto.LightChecks.Item30.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{park_tail_lights_fl}}", dto.GetCheckboxCheckStatusString(inspectDto.LightChecks.Item31.FL), -1)
			html = strings.Replace(html, "{{park_tail_lights_fr}}", dto.GetCheckboxCheckStatusString(inspectDto.LightChecks.Item31.FR), -1)
			html = strings.Replace(html, "{{park_tail_lights_rl}}", dto.GetCheckboxCheckStatusString(inspectDto.LightChecks.Item31.RL), -1)
			html = strings.Replace(html, "{{park_tail_lights_rr}}", dto.GetCheckboxCheckStatusString(inspectDto.LightChecks.Item31.RR), -1)
			html = strings.Replace(html, "{{park_tail_lights}}", inspectDto.LightChecks.Item31.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{turn_signals_rate}}", inspectDto.LightChecks.Item32.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{signal_cancellation_l}}", dto.GetCheckboxCheckStatusString(inspectDto.LightChecks.Item33.L), -1)
			html = strings.Replace(html, "{{signal_cancellation_r}}", dto.GetCheckboxCheckStatusString(inspectDto.LightChecks.Item33.R), -1)
			html = strings.Replace(html, "{{signal_cancellation}}", inspectDto.LightChecks.Item33.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{brake_lights_l}}", dto.GetCheckboxCheckStatusString(inspectDto.LightChecks.Item34.L), -1)
			html = strings.Replace(html, "{{brake_lights_r}}", dto.GetCheckboxCheckStatusString(inspectDto.LightChecks.Item34.R), -1)
			html = strings.Replace(html, "{{brake_lights_high_level}}", dto.GetCheckboxCheckStatusString(inspectDto.LightChecks.Item34.HighLevel), -1)
			html = strings.Replace(html, "{{brake_lights}}", inspectDto.LightChecks.Item34.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{license_plate_lights_l}}", dto.GetCheckboxCheckStatusString(inspectDto.LightChecks.Item35.L), -1)
			html = strings.Replace(html, "{{license_plate_lights_r}}", dto.GetCheckboxCheckStatusString(inspectDto.LightChecks.Item35.R), -1)
			html = strings.Replace(html, "{{license_plate_lights}}", inspectDto.LightChecks.Item35.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{reverse_lights_l}}", dto.GetCheckboxCheckStatusString(inspectDto.LightChecks.Item36.L), -1)
			html = strings.Replace(html, "{{reverse_lights_r}}", dto.GetCheckboxCheckStatusString(inspectDto.LightChecks.Item36.R), -1)
			html = strings.Replace(html, "{{reverse_lights}}", inspectDto.LightChecks.Item36.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{all_lenses_condition}}", inspectDto.LightChecks.Item37.GetPassOrFailString(), -1)
		}

		// 4. Interior Checks
		if inspectDto.InteriorChecks != nil {
			html = strings.Replace(html, "{{instrument_warning_lights}}", inspectDto.InteriorChecks.Item38.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{engine_check_light}}", inspectDto.InteriorChecks.Item39.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{instrument_panel_lights}}", inspectDto.InteriorChecks.Item40.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{interior_lights_courtesy_lights}}", inspectDto.InteriorChecks.Item41.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{windscreen_wiper_blades_l}}", dto.GetCheckboxCheckStatusString(inspectDto.InteriorChecks.Item42.L), -1)
			html = strings.Replace(html, "{{windscreen_wiper_blades_r}}", dto.GetCheckboxCheckStatusString(inspectDto.InteriorChecks.Item42.R), -1)
			html = strings.Replace(html, "{{windscreen_wiper_blades_rear}}", dto.GetCheckboxCheckStatusString(inspectDto.InteriorChecks.Item42.Rear), -1)
			html = strings.Replace(html, "{{windscreen_wiper_blades}}", inspectDto.InteriorChecks.Item42.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{windscreen_condition_and_visibility}}", inspectDto.InteriorChecks.Item43.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{mirrors_internal_external}}", inspectDto.InteriorChecks.Item44.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{horn_operation}}", inspectDto.InteriorChecks.Item45.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{seat_belts_fl}}", dto.GetCheckboxCheckStatusString(inspectDto.InteriorChecks.Item46.FL), -1)
			html = strings.Replace(html, "{{seat_belts_fr}}", dto.GetCheckboxCheckStatusString(inspectDto.InteriorChecks.Item46.FR), -1)
			html = strings.Replace(html, "{{seat_belts_rl}}", dto.GetCheckboxCheckStatusString(inspectDto.InteriorChecks.Item46.RL), -1)
			html = strings.Replace(html, "{{seat_belts_rc}}", dto.GetCheckboxCheckStatusString(inspectDto.InteriorChecks.Item46.RC), -1)
			html = strings.Replace(html, "{{seat_belts_rr}}", dto.GetCheckboxCheckStatusString(inspectDto.InteriorChecks.Item46.RR), -1)
			html = strings.Replace(html, "{{seat_belts}}", inspectDto.InteriorChecks.Item46.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{lubricate_door_locks_check_straps_hinges_bonnet_latch_fl}}", dto.GetCheckboxCheckStatusString(inspectDto.InteriorChecks.Item47.FL), -1)
			html = strings.Replace(html, "{{lubricate_door_locks_check_straps_hinges_bonnet_latch_fr}}", dto.GetCheckboxCheckStatusString(inspectDto.InteriorChecks.Item47.FR), -1)
			html = strings.Replace(html, "{{lubricate_door_locks_check_straps_hinges_bonnet_latch_rl}}", dto.GetCheckboxCheckStatusString(inspectDto.InteriorChecks.Item47.RL), -1)
			html = strings.Replace(html, "{{lubricate_door_locks_check_straps_hinges_bonnet_latch_rr}}", dto.GetCheckboxCheckStatusString(inspectDto.InteriorChecks.Item47.RR), -1)
			html = strings.Replace(html, "{{lubricate_door_locks_check_straps_hinges_bonnet_latch_boot}}", dto.GetCheckboxCheckStatusString(inspectDto.InteriorChecks.Item47.Boot), -1)
			html = strings.Replace(html, "{{lubricate_door_locks_check_straps_hinges_bonnet_latch}}", inspectDto.InteriorChecks.Item47.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{window_operation}}", inspectDto.InteriorChecks.Item48.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{boot_and_tail_gate_operation_fl}}", dto.GetCheckboxCheckStatusString(inspectDto.InteriorChecks.Item49.FL), -1)
			html = strings.Replace(html, "{{boot_and_tail_gate_operation_fr}}", dto.GetCheckboxCheckStatusString(inspectDto.InteriorChecks.Item49.FR), -1)
			html = strings.Replace(html, "{{boot_and_tail_gate_operation_rl}}", dto.GetCheckboxCheckStatusString(inspectDto.InteriorChecks.Item49.RL), -1)
			html = strings.Replace(html, "{{boot_and_tail_gate_operation_rc}}", dto.GetCheckboxCheckStatusString(inspectDto.InteriorChecks.Item49.RC), -1)
			html = strings.Replace(html, "{{boot_and_tail_gate_operation}}", inspectDto.InteriorChecks.Item49.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{pollen_filter}}", inspectDto.InteriorChecks.Item50.GetPassOrFailString(), -1)
		}

		// 5. Under Body
		if inspectDto.UnderBody != nil {
			html = strings.Replace(html, "{{drain_oil_replace_sump_plug_washer}}", inspectDto.UnderBody.Item51.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{replace_oil_filter}}", inspectDto.UnderBody.Item52.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{engine_oil_leaks}}", inspectDto.UnderBody.Item53.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{gearbox_oil_level_leaks}}", inspectDto.UnderBody.Item54.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{differential_oil_level_leaks}}", inspectDto.UnderBody.Item55.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{lubricate_suspension_where_applicable}}", inspectDto.UnderBody.Item56.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{fuel_line_leaks_attachments}}", inspectDto.UnderBody.Item57.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{brake_cables_hoses_lines}}", inspectDto.UnderBody.Item58.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{engine_mountings_gearbox_mounts}}", dto.GetCheckboxCheckStatusString(inspectDto.UnderBody.Item59.GearboxMounts), -1)
			html = strings.Replace(html, "{{engine_mountings_front}}", dto.GetCheckboxCheckStatusString(inspectDto.UnderBody.Item59.Front), -1)
			html = strings.Replace(html, "{{engine_mountings_rear}}", dto.GetCheckboxCheckStatusString(inspectDto.UnderBody.Item59.Rear), -1)
			html = strings.Replace(html, "{{engine_mountings}}", inspectDto.UnderBody.Item59.GetPassOrFailString(), -1)
		}

		// 6. Exhaust System Checks
		if inspectDto.ExhaustSystemChecks != nil {
			html = strings.Replace(html, "{{engine_pipe_flange}}", inspectDto.ExhaustSystemChecks.Item60.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{mufflers_resonators_pipes}}", inspectDto.ExhaustSystemChecks.Item61.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{catalytic_converter_visual}}", inspectDto.ExhaustSystemChecks.Item62.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{particulate_filter_visual}}", inspectDto.ExhaustSystemChecks.Item63.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{support_hangers}}", inspectDto.ExhaustSystemChecks.Item64.GetPassOrFailString(), -1)
		}

		// 7. Suspension / Steering System Test
		if inspectDto.SuspensionSteeringSystemTest != nil {
			html = strings.Replace(html, "{{steering_free_play}}", inspectDto.SuspensionSteeringSystemTest.Item65.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{steering_wear_leaks_rack_boots}}", inspectDto.SuspensionSteeringSystemTest.Item66.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{tie_rod_ends}}", inspectDto.SuspensionSteeringSystemTest.Item67.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{suspension_bushes}}", inspectDto.SuspensionSteeringSystemTest.Item68.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{upper_and_lower_wishbones}}", inspectDto.SuspensionSteeringSystemTest.Item69.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{sway_bar_rubbers_front}}", dto.GetCheckboxCheckStatusString(inspectDto.SuspensionSteeringSystemTest.Item70.Front), -1)
			html = strings.Replace(html, "{{sway_bar_rubbers_rear}}", dto.GetCheckboxCheckStatusString(inspectDto.SuspensionSteeringSystemTest.Item70.Rear), -1)
			html = strings.Replace(html, "{{sway_bar_rubbers}}", inspectDto.SuspensionSteeringSystemTest.Item70.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{ball_joints}}", inspectDto.SuspensionSteeringSystemTest.Item71.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{shock_absorbers_fl}}", dto.GetCheckboxCheckStatusString(inspectDto.SuspensionSteeringSystemTest.Item72.FL), -1)
			html = strings.Replace(html, "{{shock_absorbers_fr}}", dto.GetCheckboxCheckStatusString(inspectDto.SuspensionSteeringSystemTest.Item72.FR), -1)
			html = strings.Replace(html, "{{shock_absorbers_rl}}", dto.GetCheckboxCheckStatusString(inspectDto.SuspensionSteeringSystemTest.Item72.RL), -1)
			html = strings.Replace(html, "{{shock_absorbers_rr}}", dto.GetCheckboxCheckStatusString(inspectDto.SuspensionSteeringSystemTest.Item72.RR), -1)
			html = strings.Replace(html, "{{shock_absorbers}}", inspectDto.SuspensionSteeringSystemTest.Item72.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{tail_shaft_joints}}", inspectDto.SuspensionSteeringSystemTest.Item73.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{constant_velocity_joints_fl}}", dto.GetCheckboxCheckStatusString(inspectDto.SuspensionSteeringSystemTest.Item74.FL), -1)
			html = strings.Replace(html, "{{constant_velocity_joints_fr}}", dto.GetCheckboxCheckStatusString(inspectDto.SuspensionSteeringSystemTest.Item74.FR), -1)
			html = strings.Replace(html, "{{constant_velocity_joints_rl}}", dto.GetCheckboxCheckStatusString(inspectDto.SuspensionSteeringSystemTest.Item74.RL), -1)
			html = strings.Replace(html, "{{constant_velocity_joints_rr}}", dto.GetCheckboxCheckStatusString(inspectDto.SuspensionSteeringSystemTest.Item74.RR), -1)
			html = strings.Replace(html, "{{constant_velocity_joints}}", inspectDto.SuspensionSteeringSystemTest.Item74.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{rear_spring_bushes_l}}", dto.GetCheckboxCheckStatusString(inspectDto.SuspensionSteeringSystemTest.Item75.L), -1)
			html = strings.Replace(html, "{{rear_spring_bushes_r}}", dto.GetCheckboxCheckStatusString(inspectDto.SuspensionSteeringSystemTest.Item75.R), -1)
			html = strings.Replace(html, "{{rear_spring_bushes}}", inspectDto.SuspensionSteeringSystemTest.Item75.GetPassOrFailString(), -1)
		}

		// 8. Breaking System Test
		if inspectDto.BreakingSystemTest != nil {
			html = strings.Replace(html, "{{tyre_pressure_status_fl}}", inspectDto.BreakingSystemTest.Item76.FL, -1)
			html = strings.Replace(html, "{{tyre_pressure_status_fr}}", inspectDto.BreakingSystemTest.Item76.FR, -1)
			html = strings.Replace(html, "{{tyre_pressure_status_rl}}", inspectDto.BreakingSystemTest.Item76.RL, -1)
			html = strings.Replace(html, "{{tyre_pressure_status_rr}}", inspectDto.BreakingSystemTest.Item76.RR, -1)
			html = strings.Replace(html, "{{tyre_pressure_status_sp}}", inspectDto.BreakingSystemTest.Item76.SP, -1)
			html = strings.Replace(html, "{{tyre_pressure_status}}", inspectDto.BreakingSystemTest.Item76.GetPassOrFailString(), -1)

			html = strings.Replace(html, "{{tyre_condition_status_fl}}", dto.GetCheckboxCheckStatusString(inspectDto.BreakingSystemTest.Item77.FL), -1)
			html = strings.Replace(html, "{{tyre_condition_status_fr}}", dto.GetCheckboxCheckStatusString(inspectDto.BreakingSystemTest.Item77.FR), -1)
			html = strings.Replace(html, "{{tyre_condition_status_rl}}", dto.GetCheckboxCheckStatusString(inspectDto.BreakingSystemTest.Item77.RL), -1)
			html = strings.Replace(html, "{{tyre_condition_status_rr}}", dto.GetCheckboxCheckStatusString(inspectDto.BreakingSystemTest.Item77.RR), -1)
			html = strings.Replace(html, "{{tyre_condition_status_sp}}", dto.GetCheckboxCheckStatusString(inspectDto.BreakingSystemTest.Item77.SP), -1)
			html = strings.Replace(html, "{{tyre_condition_status}}", inspectDto.BreakingSystemTest.Item77.GetPassOrFailString(), -1)

			html = strings.Replace(html, "{{wheel_bearings_status_fl}}", dto.GetCheckboxCheckStatusString(inspectDto.BreakingSystemTest.Item78.FL), -1)
			html = strings.Replace(html, "{{wheel_bearings_status_fr}}", dto.GetCheckboxCheckStatusString(inspectDto.BreakingSystemTest.Item78.FR), -1)
			html = strings.Replace(html, "{{wheel_bearings_status_rl}}", dto.GetCheckboxCheckStatusString(inspectDto.BreakingSystemTest.Item78.RL), -1)
			html = strings.Replace(html, "{{wheel_bearings_status_rr}}", dto.GetCheckboxCheckStatusString(inspectDto.BreakingSystemTest.Item78.RR), -1)
			html = strings.Replace(html, "{{wheel_bearings_status}}", inspectDto.BreakingSystemTest.Item78.GetPassOrFailString(), -1)

			html = strings.Replace(html, "{{master_cyl_booster_status_master}}", dto.GetCheckboxCheckStatusString(inspectDto.BreakingSystemTest.Item79.Master), -1)
			html = strings.Replace(html, "{{master_cyl_booster_status_booster}}", dto.GetCheckboxCheckStatusString(inspectDto.BreakingSystemTest.Item79.Booster), -1)
			html = strings.Replace(html, "{{master_cyl_booster_status}}", inspectDto.BreakingSystemTest.Item79.GetPassOrFailString(), -1)

			html = strings.Replace(html, "{{brakes_lh_front_manufacturers_spec_size}}", inspectDto.BreakingSystemTest.Item80.LH_Front.ManufacturersSpecSize, -1)
			html = strings.Replace(html, "{{brakes_lh_front_disc_size}}", inspectDto.BreakingSystemTest.Item80.LH_Front.DiscOrDrumSize, -1)
			html = strings.Replace(html, "{{brakes_lh_front_pad_worn}}", inspectDto.BreakingSystemTest.Item80.LH_Front.PadOrLiningPercentageWorn, -1)
			html = strings.Replace(html, "{{brakes_lh_front_caliper}}", inspectDto.BreakingSystemTest.Item80.LH_Front.CaliperOrCylinder.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{brakes_lh_front_seats}}", inspectDto.BreakingSystemTest.Item80.LH_Front.Seats.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{brakes_lh_front}}", inspectDto.BreakingSystemTest.Item80.LH_Front.Seats.GetPassOrFailString(), -1)

			html = strings.Replace(html, "{{brakes_rh_front_manufacturers_spec_size}}", inspectDto.BreakingSystemTest.Item80.RH_Front.ManufacturersSpecSize, -1)
			html = strings.Replace(html, "{{brakes_rh_front_disc_size}}", inspectDto.BreakingSystemTest.Item80.RH_Front.DiscOrDrumSize, -1)
			html = strings.Replace(html, "{{brakes_rh_front_pad_worn}}", inspectDto.BreakingSystemTest.Item80.RH_Front.PadOrLiningPercentageWorn, -1)
			html = strings.Replace(html, "{{brakes_rh_front_caliper}}", inspectDto.BreakingSystemTest.Item80.RH_Front.CaliperOrCylinder.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{brakes_rh_front_seats}}", inspectDto.BreakingSystemTest.Item80.RH_Front.Seats.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{brakes_rh_front}}", inspectDto.BreakingSystemTest.Item80.RH_Front.Seats.GetPassOrFailString(), -1)

			html = strings.Replace(html, "{{brakes_lh_rear_manufacturers_spec_size}}", inspectDto.BreakingSystemTest.Item80.LH_Rear.ManufacturersSpecSize, -1)
			html = strings.Replace(html, "{{brakes_lh_rear_disc_size}}", inspectDto.BreakingSystemTest.Item80.LH_Rear.DiscOrDrumSize, -1)
			html = strings.Replace(html, "{{brakes_lh_rear_pad_worn}}", inspectDto.BreakingSystemTest.Item80.LH_Rear.PadOrLiningPercentageWorn, -1)
			html = strings.Replace(html, "{{brakes_lh_rear_caliper}}", inspectDto.BreakingSystemTest.Item80.LH_Rear.CaliperOrCylinder.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{brakes_lh_rear_seats}}", inspectDto.BreakingSystemTest.Item80.LH_Rear.Seats.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{brakes_lh_rear}}", inspectDto.BreakingSystemTest.Item80.LH_Rear.Seats.GetPassOrFailString(), -1)

			html = strings.Replace(html, "{{brakes_rh_rear_manufacturers_spec_size}}", inspectDto.BreakingSystemTest.Item80.RH_Rear.ManufacturersSpecSize, -1)
			html = strings.Replace(html, "{{brakes_rh_rear_disc_size}}", inspectDto.BreakingSystemTest.Item80.RH_Rear.DiscOrDrumSize, -1)
			html = strings.Replace(html, "{{brakes_rh_rear_pad_worn}}", inspectDto.BreakingSystemTest.Item80.RH_Rear.PadOrLiningPercentageWorn, -1)
			html = strings.Replace(html, "{{brakes_rh_rear_caliper}}", inspectDto.BreakingSystemTest.Item80.RH_Rear.CaliperOrCylinder.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{brakes_rh_rear_seats}}", inspectDto.BreakingSystemTest.Item80.RH_Rear.Seats.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{brakes_rh_rear}}", inspectDto.BreakingSystemTest.Item80.RH_Rear.Seats.GetPassOrFailString(), -1)

		}

		// 9. Under The Bonnet Tests
		if inspectDto.UnderTheBonnetTests != nil {
			html = strings.Replace(html, "{{replace_engine_oil}}", inspectDto.UnderTheBonnetTests.Item81.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{battery_load_test_electrolyte}}", dto.GetCheckboxCheckStatusString(inspectDto.UnderTheBonnetTests.Item82.Electrolyte), -1)
			html = strings.Replace(html, "{{battery_load_test_clamp}}", dto.GetCheckboxCheckStatusString(inspectDto.UnderTheBonnetTests.Item82.Clamp), -1)
			html = strings.Replace(html, "{{battery_load_test_terminals}}", dto.GetCheckboxCheckStatusString(inspectDto.UnderTheBonnetTests.Item82.Terminals), -1)
			html = strings.Replace(html, "{{battery_load_test_cables}}", dto.GetCheckboxCheckStatusString(inspectDto.UnderTheBonnetTests.Item82.Cables), -1)
			html = strings.Replace(html, "{{battery_load_test}}", inspectDto.UnderTheBonnetTests.Item82.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{air_cleaner}}", inspectDto.UnderTheBonnetTests.Item83.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{belts}}", inspectDto.UnderTheBonnetTests.Item84.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{cambelt_manuf_spec}}", inspectDto.UnderTheBonnetTests.Item85.ManufSpec, -1)
			html = strings.Replace(html, "{{cambelt_date}}", inspectDto.UnderTheBonnetTests.Item85.Date, -1)
			html = strings.Replace(html, "{{cambelt}}", inspectDto.UnderTheBonnetTests.Item85.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{coolant_condition}}", inspectDto.UnderTheBonnetTests.Item86.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{external_visual_check}}", inspectDto.UnderTheBonnetTests.Item87.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{start_motor_check_oil_filter}}", inspectDto.UnderTheBonnetTests.Item88.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{power_steering_condition}}", inspectDto.UnderTheBonnetTests.Item89.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{transmission_oil_check}}", inspectDto.UnderTheBonnetTests.Item90.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{brake_clutch_fluid}}", inspectDto.UnderTheBonnetTests.Item91.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{windscreen_washers}}", inspectDto.UnderTheBonnetTests.Item92.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{fuel_filter}}", inspectDto.UnderTheBonnetTests.Item93.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{pressurise_cooling_system}}", inspectDto.UnderTheBonnetTests.Item94.GetPassOrFailString(), -1)
			html = strings.Replace(html, "{{bonnet_boot_latch}}", inspectDto.UnderTheBonnetTests.Item95.GetPassOrFailString(), -1)
		}

		// 10. Final Procedures
		if inspectDto.FinalProcedures != nil {
			html = strings.Replace(html, "{{road_test_max_speed_reached}}", inspectDto.FinalProcedures.Item96.MaxSpeed, -1)
			html = strings.Replace(html, "{{road_test_max_speed}}", inspectDto.FinalProcedures.Item96.MaxSpeed, -1)
			html = strings.Replace(html, "{{park_vehicle_facing_out}}", dto.GetCheckboxCheckStatusString(inspectDto.FinalProcedures.Item97), -1)
			html = strings.Replace(html, "{{reset_service_interval}}", dto.GetCheckboxCheckStatusString(inspectDto.FinalProcedures.Item98), -1)
			html = strings.Replace(html, "{{gloss_tyres}}", dto.GetCheckboxCheckStatusString(inspectDto.FinalProcedures.Item99), -1)
			html = strings.Replace(html, "{{vacuum_carpets_and_deodorise}}", dto.GetCheckboxCheckStatusString(inspectDto.FinalProcedures.Item100), -1)
			html = strings.Replace(html, "{{deodorise_interior}}", dto.GetCheckboxCheckStatusString(inspectDto.FinalProcedures.Item101), -1)
			html = strings.Replace(html, "{{wipe_over_dash}}", dto.GetCheckboxCheckStatusString(inspectDto.FinalProcedures.Item102), -1)
			html = strings.Replace(html, "{{clean_windows}}", dto.GetCheckboxCheckStatusString(inspectDto.FinalProcedures.Item103), -1)
			html = strings.Replace(html, "{{remove_seat_cover_and_floor_mat}}", dto.GetCheckboxCheckStatusString(inspectDto.FinalProcedures.Item104), -1)
		}
	}

	return html
}
