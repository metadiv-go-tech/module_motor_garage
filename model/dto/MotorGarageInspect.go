package dto

import "encoding/json"

type MotorGarageInspect struct {
	ID uint `json:"id"`

	InvoiceId uint                `json:"invoice_id"`
	Invoice   *MotorGarageInvoice `json:"invoice"`

	RoadTest *MotorGarageInspectRoadTest `json:"road_test"`
}

type MotorGarageInspectPassOrFail struct {
	Pass bool `json:"pass"`
}

type MotorGarageInspectItem4 struct {
	MotorGarageInspectPassOrFail
	Speed string `json:"speed"`
}

type MotorGarageInspectRoadTest struct {
	Item1 *MotorGarageInspectPassOrFail `json:"item_1"`
	Item4 *MotorGarageInspectItem4      `json:"item_4"`
}

func (d *MotorGarageInspectRoadTest) ToString() string {
	b, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	return string(b)
}
