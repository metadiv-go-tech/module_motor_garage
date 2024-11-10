package request

import (
	"errors"

	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

type MotorGarageTechnicianCreate struct {
	Name string `json:"name"`
}

type MotorGarageTechnicianUpdate struct {
	base.RequestPathId
	MotorGarageTechnicianCreate
}

func (r *MotorGarageTechnicianCreate) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

func (r *MotorGarageTechnicianCreate) ToEntity(e *entity.MotorGarageTechnician) *entity.MotorGarageTechnician {
	if e == nil {
		e = &entity.MotorGarageTechnician{}
	}
	e.Name = r.Name
	return e
}
