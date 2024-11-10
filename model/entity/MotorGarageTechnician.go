package entity

import (
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
)

type MotorGarageTechnician struct {
	base.Model
	base.ModelWorkspace

	Name string `json:"name"`
}

func (e *MotorGarageTechnician) ToDTO() *dto.MotorGarageTechnician {
	return &dto.MotorGarageTechnician{
		ID:   e.ID,
		Name: e.Name,
	}
}
