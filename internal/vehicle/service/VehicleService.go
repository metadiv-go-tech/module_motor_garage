package service

import (
	"github.com/metadiv-go-tech/metaorm"
	"github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
	"gorm.io/gorm"
)

var VehicleService = new(vehicleService)

type vehicleService struct{}

func (s *vehicleService) CheckRegistration(tx *gorm.DB, registration string, workspaceId uint, excludeId uint) (duplicated bool) {
	exist := repo.MotorGarageVehicleRepo.FindOne(tx,
		metaorm.And(
			metaorm.Eq("registration", registration),
			metaorm.Neq("id", excludeId),
		),
		workspaceId,
	)
	return exist != nil
}
