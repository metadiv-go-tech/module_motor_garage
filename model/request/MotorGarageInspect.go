package request

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

type MotorGarageInspectCreate struct {
	InvoiceId uint              `json:"invoice_id"`
	PassItems []string          `json:"pass_items"`
	FailItems []string          `json:"fail_items"`
	KeyValues map[string]string `json:"key_values"`
}

type MotorGarageInspectUpdate struct {
	base.RequestPathId
	MotorGarageInspectCreate
}

func (r *MotorGarageInspectCreate) Validate() error {
	if r.InvoiceId == 0 {
		return errors.New("invoice is required")
	}
	return nil
}

func (r *MotorGarageInspectCreate) ToEntity(e *entity.MotorGarageInspect) *entity.MotorGarageInspect {
	if e == nil {
		e = &entity.MotorGarageInspect{}
	}
	e.InvoiceId = r.InvoiceId
	if len(r.PassItems) > 0 {
		e.PassItems = strings.Join(r.PassItems, ",")
	}
	if len(r.FailItems) > 0 {
		e.FailItems = strings.Join(r.FailItems, ",")
	}
	if len(r.KeyValues) > 0 {
		bytes, _ := json.Marshal(r.KeyValues)
		e.KeyValues = string(bytes)
	}
	return e
}
