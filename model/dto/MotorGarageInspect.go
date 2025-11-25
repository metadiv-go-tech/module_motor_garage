package dto

import "encoding/json"

type MotorGarageInspect struct {
	ID uint `json:"id"`

	InvoiceId uint                `json:"invoice_id"`
	Invoice   *MotorGarageInvoice `json:"invoice"`

	// Inspect Sections
	CustomerInstructionsAndRepairs *MotorGarageCustomerInstructionsAndRepairs      `json:"customer_instructions_and_repairs"`
	RoadTest                       *MotorGarageInspectRoadTest                     `json:"road_test"`
	EngineTune                     *MotorGarageInspectEngineTune                   `json:"engine_tune"`
	LightChecks                    *MotorGarageInspectLightChecks                  `json:"light_checks"`
	InteriorChecks                 *MotorGarageInspectInteriorChecks               `json:"interior_checks"`
	UnderBody                      *MotorGarageInspectUnderBody                    `json:"under_body"`
	ExhaustSystemChecks            *MotorGarageInspectExhaustSystemChecks          `json:"exhaust_system_checks"`
	SuspensionSteeringSystemTest   *MotorGarageInspectSuspensionSteeringSystemTest `json:"suspension_steering_system_test"`
	BreakingSystemTest             *MotorGarageInspectBreakingSystemTest           `json:"breaking_system_test"`
	UnderTheBonnetTests            *MotorGarageInspectUnderTheBonnetTests          `json:"under_the_bonnet_tests"`
	FinalProcedures                *MotorGarageInspectFinalProcedures              `json:"final_procedures"`
}

// Inspect Sections general items
type MotorGarageInspectPassOrFail struct {
	Pass *bool `json:"pass"` // 使用指针以支持 null 值，nil = N/A, &true = OK, &false = FAIL
}

// MarshalJSON 确保即使 Pass 为 nil 时也输出 pass: null
// 这样即使 frontend 不传 pass 字段，返回时也会显示 pass: null
func (p *MotorGarageInspectPassOrFail) MarshalJSON() ([]byte, error) {
	if p == nil {
		return json.Marshal(map[string]interface{}{
			"pass": nil,
		})
	}
	// 明确输出 Pass 字段，即使为 nil 也会序列化为 null
	return json.Marshal(map[string]interface{}{
		"pass": p.Pass,
	})
}

func (p *MotorGarageInspectPassOrFail) GetPassOrFailString() string {
	if p == nil || p.Pass == nil {
		return "N/A"
	}

	if *p.Pass {
		return "OK"
	}

	return "FAIL"
}

type MotorGarageInspectLeftRightOptions struct {
	L bool `json:"l"`
	R bool `json:"r"`
}

type MotorGarageInspectDirectionOptions struct {
	FL bool `json:"fl"`
	FR bool `json:"fr"`
	RL bool `json:"rl"`
	RR bool `json:"rr"`
}

type MotorGarageInspectPassOrFailWithLeftRightOptions struct {
	*MotorGarageInspectPassOrFail
	MotorGarageInspectLeftRightOptions
}

type MotorGarageInspectPassOrFailWithDirectionOptions struct {
	*MotorGarageInspectPassOrFail
	MotorGarageInspectDirectionOptions
}

// Customer Instructions and Repairs
type MotorGarageCustomerInstructionsAndRepairs struct {
	Instruction1       bool   `json:"instruction_1"`
	Instruction2       bool   `json:"instruction_2"`
	Instruction3       bool   `json:"instruction_3"`
	Instruction4       bool   `json:"instruction_4"`
	Instruction5       bool   `json:"instruction_5"`
	Instruction6       bool   `json:"instruction_6"`
	Instruction7       bool   `json:"instruction_7"`
	Instruction8       bool   `json:"instruction_8"`
	Instruction9       bool   `json:"instruction_9"`
	Instruction10      bool   `json:"instruction_10"`
	Instruction11      bool   `json:"instruction_11"`
	Instruction12      bool   `json:"instruction_12"`
	Instruction13      bool   `json:"instruction_13"`
	Instruction14      bool   `json:"instruction_14"`
	Instruction15      bool   `json:"instruction_15"`
	Instruction16      bool   `json:"instruction_16"`
	PrimeItemOfConcern string `json:"prime_item_of_concern"`
}

// 1. Road Test
type MotorGarageInspectItem4 struct {
	*MotorGarageInspectPassOrFail
	Speed string `json:"speed"`
}

type MotorGarageInspectRoadTest struct {
	Item1  *MotorGarageInspectPassOrFail `json:"item_1"`
	Item2  *MotorGarageInspectPassOrFail `json:"item_2"`
	Item3  *MotorGarageInspectPassOrFail `json:"item_3"`
	Item4  MotorGarageInspectItem4       `json:"item_4"`
	Item5  *MotorGarageInspectPassOrFail `json:"item_5"`
	Item6  *MotorGarageInspectPassOrFail `json:"item_6"`
	Item7  *MotorGarageInspectPassOrFail `json:"item_7"`
	Item8  *MotorGarageInspectPassOrFail `json:"item_8"`
	Item9  *MotorGarageInspectPassOrFail `json:"item_9"`
	Item10 *MotorGarageInspectPassOrFail `json:"item_10"`
	Item11 *MotorGarageInspectPassOrFail `json:"item_11"`
	Item12 *MotorGarageInspectPassOrFail `json:"item_12"`
}

// 2. Engine Tune
type MotorGarageInspectBeforeAfter struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

func (b *MotorGarageInspectBeforeAfter) GetBeforeAfterString() (string, string) {
	if b == nil {
		return "", ""
	}
	return b.Before, b.After
}

func GetCheckboxCheckStatusString(checked bool) string {
	if checked {
		return "checked"
	}

	return ""
}

type MotorGarageInspectItem26 struct {
	MotorGarageInspectBeforeAfter
	Levels bool `json:"levels"`
}

type MotorGarageInspectEngineTune struct {
	Item13 MotorGarageInspectBeforeAfter `json:"item_13"`
	Item14 MotorGarageInspectBeforeAfter `json:"item_14"`
	Item15 MotorGarageInspectBeforeAfter `json:"item_15"`
	Item16 MotorGarageInspectBeforeAfter `json:"item_16"`
	Item17 MotorGarageInspectBeforeAfter `json:"item_17"`
	Item18 MotorGarageInspectBeforeAfter `json:"item_18"`
	Item19 MotorGarageInspectBeforeAfter `json:"item_19"`
	Item20 MotorGarageInspectBeforeAfter `json:"item_20"`
	Item21 MotorGarageInspectBeforeAfter `json:"item_21"`
	Item22 MotorGarageInspectBeforeAfter `json:"item_22"`
	Item23 MotorGarageInspectBeforeAfter `json:"item_23"`
	Item24 MotorGarageInspectBeforeAfter `json:"item_24"`
	Item25 MotorGarageInspectBeforeAfter `json:"item_25"`
	Item26 MotorGarageInspectItem26      `json:"item_26"`
	Item27 MotorGarageInspectBeforeAfter `json:"item_27"`
	Item28 []string                      `json:"item_28"`
}

// 3. Light checks
type MotorGarageInspectItem29 struct {
	*MotorGarageInspectPassOrFail
	Hi_L bool `json:"hi_l"`
	Hi_R bool `json:"hi_r"`
	Lo_L bool `json:"lo_l"`
	Lo_R bool `json:"lo_r"`
}

type MotorGarageInspectItem34 struct {
	MotorGarageInspectPassOrFailWithLeftRightOptions
	HighLevel bool `json:"high_level"`
}

type MotorGarageInspectLightChecks struct {
	Item29 MotorGarageInspectItem29                         `json:"item_29"`
	Item30 *MotorGarageInspectPassOrFail                    `json:"item_30"`
	Item31 MotorGarageInspectPassOrFailWithDirectionOptions `json:"item_31"`
	Item32 *MotorGarageInspectPassOrFail                    `json:"item_32"`
	Item33 MotorGarageInspectPassOrFailWithLeftRightOptions `json:"item_33"`
	Item34 MotorGarageInspectItem34                         `json:"item_34"`
	Item35 MotorGarageInspectPassOrFailWithLeftRightOptions `json:"item_35"`
	Item36 MotorGarageInspectPassOrFailWithLeftRightOptions `json:"item_36"`
	Item37 *MotorGarageInspectPassOrFail                    `json:"item_37"`
}

// 4. Interior Checks
type MotorGarageInspectItem42 struct {
	MotorGarageInspectPassOrFailWithLeftRightOptions
	Rear bool `json:"rear"`
}

type MotorGarageInspectItem46 struct {
	MotorGarageInspectPassOrFailWithDirectionOptions
	RC bool `json:"rc"`
}

type MotorGarageInspectItem47 struct {
	MotorGarageInspectPassOrFailWithDirectionOptions
	Boot bool `json:"boot"`
}

type MotorGarageInspectInteriorChecks struct {
	Item38 *MotorGarageInspectPassOrFail                    `json:"item_38"`
	Item39 *MotorGarageInspectPassOrFail                    `json:"item_39"`
	Item40 *MotorGarageInspectPassOrFail                    `json:"item_40"`
	Item41 *MotorGarageInspectPassOrFail                    `json:"item_41"`
	Item42 MotorGarageInspectItem42                         `json:"item_42"`
	Item43 *MotorGarageInspectPassOrFail                    `json:"item_43"`
	Item44 *MotorGarageInspectPassOrFail                    `json:"item_44"`
	Item45 *MotorGarageInspectPassOrFail                    `json:"item_45"`
	Item46 MotorGarageInspectItem46                         `json:"item_46"`
	Item47 MotorGarageInspectItem47                         `json:"item_47"`
	Item48 MotorGarageInspectPassOrFailWithDirectionOptions `json:"item_48"`
	Item49 *MotorGarageInspectPassOrFail                    `json:"item_49"`
	Item50 *MotorGarageInspectPassOrFail                    `json:"item_50"`
}

// 5. Under Body
type MotorGarageInspectItem59 struct {
	*MotorGarageInspectPassOrFail
	GearboxMounts bool `json:"gearbox_mounts"`
	Front         bool `json:"front"`
	Rear          bool `json:"rear"`
}

type MotorGarageInspectUnderBody struct {
	Item51 *MotorGarageInspectPassOrFail `json:"item_51"`
	Item52 *MotorGarageInspectPassOrFail `json:"item_52"`
	Item53 *MotorGarageInspectPassOrFail `json:"item_53"`
	Item54 *MotorGarageInspectPassOrFail `json:"item_54"`
	Item55 *MotorGarageInspectPassOrFail `json:"item_55"`
	Item56 *MotorGarageInspectPassOrFail `json:"item_56"`
	Item57 *MotorGarageInspectPassOrFail `json:"item_57"`
	Item58 *MotorGarageInspectPassOrFail `json:"item_58"`
	Item59 MotorGarageInspectItem59      `json:"item_59"`
}

// 6. Exhaust System Checks
type MotorGarageInspectExhaustSystemChecks struct {
	Item60 *MotorGarageInspectPassOrFail `json:"item_60"`
	Item61 *MotorGarageInspectPassOrFail `json:"item_61"`
	Item62 *MotorGarageInspectPassOrFail `json:"item_62"`
	Item63 *MotorGarageInspectPassOrFail `json:"item_63"`
	Item64 *MotorGarageInspectPassOrFail `json:"item_64"`
}

// 7. Suspension / Steering System Test
type MotorGarageInspectItem70 struct {
	*MotorGarageInspectPassOrFail
	Front bool `json:"front"`
	Rear  bool `json:"rear"`
}

type MotorGarageInspectSuspensionSteeringSystemTest struct {
	Item65 *MotorGarageInspectPassOrFail                    `json:"item_65"`
	Item66 *MotorGarageInspectPassOrFail                    `json:"item_66"`
	Item67 *MotorGarageInspectPassOrFail                    `json:"item_67"`
	Item68 *MotorGarageInspectPassOrFail                    `json:"item_68"`
	Item69 *MotorGarageInspectPassOrFail                    `json:"item_69"`
	Item70 MotorGarageInspectItem70                         `json:"item_70"`
	Item71 *MotorGarageInspectPassOrFail                    `json:"item_71"`
	Item72 MotorGarageInspectPassOrFailWithDirectionOptions `json:"item_72"`
	Item73 *MotorGarageInspectPassOrFail                    `json:"item_73"`
	Item74 MotorGarageInspectPassOrFailWithDirectionOptions `json:"item_74"`
	Item75 MotorGarageInspectPassOrFailWithLeftRightOptions `json:"item_75"`
}

// 8. Breaking System Test
type MotorGarageInspectItem76 struct {
	*MotorGarageInspectPassOrFail
	FL string `json:"fl"`
	FR string `json:"fr"`
	RL string `json:"rl"`
	RR string `json:"rr"`
	SP string `json:"sp"`
}

type MotorGarageInspectItem77 struct {
	MotorGarageInspectPassOrFailWithDirectionOptions
	SP bool `json:"sp"`
}

type MotorGarageInspectItem79 struct {
	*MotorGarageInspectPassOrFail
	Master  bool `json:"master"`
	Booster bool `json:"booster"`
}

type BrakesTestingByDirectionValues struct {
	*MotorGarageInspectPassOrFail
	ManufacturersSpecSize     string                        `json:"manufacturers_spec_size"`
	DiscOrDrumSize            string                        `json:"disc_or_drum_size"`
	PadOrLiningPercentageWorn string                        `json:"pad_or_lining_percentage_worn"`
	CaliperOrCylinder         *MotorGarageInspectPassOrFail `json:"caliper_or_cylinder"`
	Seats                     *MotorGarageInspectPassOrFail `json:"seats"`
}

type MotorGarageInspectItem80 struct {
	LH_Front BrakesTestingByDirectionValues `json:"lh_front"`
	RH_Front BrakesTestingByDirectionValues `json:"rh_front"`
	LH_Rear  BrakesTestingByDirectionValues `json:"lh_rear"`
	RH_Rear  BrakesTestingByDirectionValues `json:"rh_rear"`
}

type MotorGarageInspectBreakingSystemTest struct {
	Item76 MotorGarageInspectItem76                         `json:"item_76"`
	Item77 MotorGarageInspectItem77                         `json:"item_77"`
	Item78 MotorGarageInspectPassOrFailWithDirectionOptions `json:"item_78"`
	Item79 MotorGarageInspectItem79                         `json:"item_79"`
	Item80 MotorGarageInspectItem80                         `json:"item_80"`
}

// 9. Under The Bonnet Tests
type MotorGarageInspectItem82 struct {
	*MotorGarageInspectPassOrFail
	Electrolyte bool `json:"electrolyte"`
	Clamp       bool `json:"clamp"`
	Terminals   bool `json:"terminals"`
	Cables      bool `json:"cables"`
}

type MotorGarageInspectItem85 struct {
	*MotorGarageInspectPassOrFail
	ManufSpec string `json:"manuf_spec"`
	Date      string `json:"date"`
}

type MotorGarageInspectUnderTheBonnetTests struct {
	Item81 *MotorGarageInspectPassOrFail `json:"item_81"`
	Item82 MotorGarageInspectItem82      `json:"item_82"`
	Item83 *MotorGarageInspectPassOrFail `json:"item_83"`
	Item84 *MotorGarageInspectPassOrFail `json:"item_84"`
	Item85 MotorGarageInspectItem85      `json:"item_85"`
	Item86 *MotorGarageInspectPassOrFail `json:"item_86"`
	Item87 *MotorGarageInspectPassOrFail `json:"item_87"`
	Item88 *MotorGarageInspectPassOrFail `json:"item_88"`
	Item89 *MotorGarageInspectPassOrFail `json:"item_89"`
	Item90 *MotorGarageInspectPassOrFail `json:"item_90"`
	Item91 *MotorGarageInspectPassOrFail `json:"item_91"`
	Item92 *MotorGarageInspectPassOrFail `json:"item_92"`
	Item93 *MotorGarageInspectPassOrFail `json:"item_93"`
	Item94 *MotorGarageInspectPassOrFail `json:"item_94"`
	Item95 *MotorGarageInspectPassOrFail `json:"item_95"`
}

// 10 .final procedures
type MotorGarageInspectItem96 struct {
	*MotorGarageInspectPassOrFail
	MaxSpeed string `json:"max_speed"`
}

type MotorGarageInspectFinalProcedures struct {
	Item96  MotorGarageInspectItem96 `json:"item_96"`
	Item97  bool                     `json:"item_97"`
	Item98  bool                     `json:"item_98"`
	Item99  bool                     `json:"item_99"`
	Item100 bool                     `json:"item_100"`
	Item101 bool                     `json:"item_101"`
	Item102 bool                     `json:"item_102"`
	Item103 bool                     `json:"item_103"`
	Item104 bool                     `json:"item_104"`
}

// helper: convert the sections to string
func InspectSectionToString[T any](d T) string {
	b, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	return string(b)
}
