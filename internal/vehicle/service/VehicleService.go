package service

import (
	"github.com/metadiv-go-tech/metaorm/v2"
	"github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/repo"
)

var VehicleService = new(vehicleService)

type vehicleService struct{}

func (s *vehicleService) CheckRegistration(db metaorm.DB, registration string, workspaceId uint, excludeId uint) (duplicated bool) {
	exist := repo.MotorGarageVehicleRepo.FindOne(db,
		db.And(
			db.DecryptedEq("registration", registration),
			db.Neq("id", excludeId),
		),
		workspaceId,
	)
	return exist != nil
}
