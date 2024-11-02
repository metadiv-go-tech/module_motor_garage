package request

import "github.com/metadiv-go-tech/metagin/v2/base"

type MotorGarageServiceUpdate struct {
	base.RequestPathId
	MotorGarageServiceCreate
}
