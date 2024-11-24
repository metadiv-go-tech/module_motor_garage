package entity

import (
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

type MotorGarageStatistic struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

func (e *MotorGarageStatistic) ToDTO(locale string) *dto.MotorGarageStatistic {
	d := &dto.MotorGarageStatistic{
		Label: e.Label,
		Value: e.Value,
	}
	return d
}
