package request

import "github.com/metadiv-go-tech/metagin/base"

type MotorGarageServiceUpdate struct {
	base.RequestIDPath
	MotorGarageServiceCreate
}
