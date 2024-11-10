package entity

import (
	"encoding/json"
	"strings"

	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

type MotorGarageInspect struct {
	base.Model
	base.ModelWorkspace

	InvoiceId uint                `json:"invoice_id"`
	Invoice   *MotorGarageInvoice `json:"invoice" gorm:"foreignKey:InvoiceId"`

	PassItems string `json:"pass_items"` // list of tags, separated by comma
	FailItems string `json:"fail_items"` // list of tags, separated by comma
	KeyValues string `json:"key_values"` // json string, tag:value pairs
}

func (e *MotorGarageInspect) ToDTO(locale string) *dto.MotorGarageInspect {
	d := &dto.MotorGarageInspect{
		ID:        e.ID,
		InvoiceId: e.InvoiceId,
	}
	if e.Invoice != nil {
		d.Invoice = e.Invoice.ToDTO(locale)
	}
	if e.PassItems != "" {
		d.PassItems = strings.Split(e.PassItems, ",")
	}
	if e.FailItems != "" {
		d.FailItems = strings.Split(e.FailItems, ",")
	}
	if e.KeyValues != "" {
		json.Unmarshal([]byte(e.KeyValues), &d.KeyValues)
	}
	return d
}
