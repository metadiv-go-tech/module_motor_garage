package dto

import "encoding/json"

type MotorGarageInspect struct {
	ID uint `json:"id"`

	InvoiceId uint                `json:"invoice_id"`
	Invoice   *MotorGarageInvoice `json:"invoice"`

	// Inspect Sections
	RoadTest       *MotorGarageInspectRoadTest       `json:"road_test"`
	EngineTune     *MotorGarageInspectEngineTune     `json:"engine_tune"`
	LightChecks    *MotorGarageInspectLightChecks    `json:"light_checks"`
	InteriorChecks *MotorGarageInspectInteriorChecks `json:"interior_checks"`
}

type MotorGarageInspectPassOrFail struct {
	Pass bool `json:"pass"`
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
	MotorGarageInspectPassOrFail
	MotorGarageInspectLeftRightOptions
}

type MotorGarageInspectPassOrFailWithDirectionOptions struct {
	MotorGarageInspectPassOrFail
	MotorGarageInspectDirectionOptions
}

// Road Test
type MotorGarageInspectItem4 struct {
	MotorGarageInspectPassOrFail
	Speed string `json:"speed"`
}

type MotorGarageInspectRoadTest struct {
	Item1  *MotorGarageInspectPassOrFail `json:"item_1"`
	Item2  *MotorGarageInspectPassOrFail `json:"item_2"`
	Item3  *MotorGarageInspectPassOrFail `json:"item_3"`
	Item4  *MotorGarageInspectItem4      `json:"item_4"`
	Item5  *MotorGarageInspectPassOrFail `json:"item_5"`
	Item6  *MotorGarageInspectPassOrFail `json:"item_6"`
	Item7  *MotorGarageInspectPassOrFail `json:"item_7"`
	Item8  *MotorGarageInspectPassOrFail `json:"item_8"`
	Item9  *MotorGarageInspectPassOrFail `json:"item_9"`
	Item10 *MotorGarageInspectPassOrFail `json:"item_10"`
	Item11 *MotorGarageInspectPassOrFail `json:"item_11"`
	Item12 *MotorGarageInspectPassOrFail `json:"item_12"`
}

// Engine Tune
type MotorGarageInspectBeforeAfter struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

type MotorGarageInspectItem26 struct {
	MotorGarageInspectBeforeAfter
	Levels bool `json:"levels"`
}

type MotorGarageInspectEngineTune struct {
	Item13 *MotorGarageInspectBeforeAfter `json:"item_13"`
	Item14 *MotorGarageInspectBeforeAfter `json:"item_14"`
	Item15 *MotorGarageInspectBeforeAfter `json:"item_15"`
	Item16 *MotorGarageInspectBeforeAfter `json:"item_16"`
	Item17 *MotorGarageInspectBeforeAfter `json:"item_17"`
	Item18 *MotorGarageInspectBeforeAfter `json:"item_18"`
	Item19 *MotorGarageInspectBeforeAfter `json:"item_19"`
	Item20 *MotorGarageInspectBeforeAfter `json:"item_20"`
	Item21 *MotorGarageInspectBeforeAfter `json:"item_21"`
	Item22 *MotorGarageInspectBeforeAfter `json:"item_22"`
	Item23 *MotorGarageInspectBeforeAfter `json:"item_23"`
	Item24 *MotorGarageInspectBeforeAfter `json:"item_24"`
	Item25 *MotorGarageInspectBeforeAfter `json:"item_25"`
	Item26 *MotorGarageInspectItem26      `json:"item_26"`
	Item27 *MotorGarageInspectBeforeAfter `json:"item_27"`
	Item28 []string                       `json:"item_28"`
}

// Light checks
type MotorGarageInspectItem29 struct {
	MotorGarageInspectPassOrFail
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
	Item29 *MotorGarageInspectItem29                         `json:"item_29"`
	Item30 *MotorGarageInspectPassOrFail                     `json:"item_30"`
	Item31 *MotorGarageInspectPassOrFailWithDirectionOptions `json:"item_31"`
	Item32 *MotorGarageInspectPassOrFail                     `json:"item_32"`
	Item33 *MotorGarageInspectPassOrFailWithLeftRightOptions `json:"item_33"`
	Item34 *MotorGarageInspectItem34                         `json:"item_34"`
	Item35 *MotorGarageInspectPassOrFailWithLeftRightOptions `json:"item_35"`
	Item36 *MotorGarageInspectPassOrFailWithLeftRightOptions `json:"item_36"`
	Item37 *MotorGarageInspectPassOrFail                     `json:"item_37"`
}

// interior checks
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
type MotorGarageInspectItem49 struct {
	MotorGarageInspectPassOrFailWithDirectionOptions
	FL bool `json:"fl"`
	FR bool `json:"fr"`
	RL bool `json:"rl"`
	RC bool `json:"rc"`
}

type MotorGarageInspectInteriorChecks struct {
	Item38 *MotorGarageInspectPassOrFail `json:"item_38"`
	Item39 *MotorGarageInspectPassOrFail `json:"item_39"`
	Item40 *MotorGarageInspectPassOrFail `json:"item_40"`
	Item41 *MotorGarageInspectPassOrFail `json:"item_41"`
	Item42 *MotorGarageInspectItem42     `json:"item_42"`
	Item43 *MotorGarageInspectPassOrFail `json:"item_43"`
	Item44 *MotorGarageInspectPassOrFail `json:"item_44"`
	Item45 *MotorGarageInspectPassOrFail `json:"item_45"`
	Item46 *MotorGarageInspectItem46     `json:"item_46"`
	Item47 *MotorGarageInspectItem47     `json:"item_47"`
	Item48 *MotorGarageInspectPassOrFail `json:"item_48"`
	Item49 *MotorGarageInspectItem49     `json:"item_49"`
	Item50 *MotorGarageInspectPassOrFail `json:"item_50"`
}

// final procedures
type MotorGarageInspectFinalProcedures struct {
	Item96  bool `json:"item_96"`
	Item97  bool `json:"item_97"`
	Item98  bool `json:"item_98"`
	Item99  bool `json:"item_99"`
	Item100 bool `json:"item_100"`
	Item101 bool `json:"item_101"`
	Item102 bool `json:"item_102"`
	Item103 bool `json:"item_103"`
	Item104 bool `json:"item_104"`
}

func InspectSectionToString[T any](d T) string {
	b, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	return string(b)
}
