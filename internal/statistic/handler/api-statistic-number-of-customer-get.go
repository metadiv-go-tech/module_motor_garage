package handler

import (
	"fmt"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"

	"github.com/metadiv-go-tech/module_motor_garage/model/dto"
	relationship "github.com/metadiv-go-tech/module_relationship/v2"
)

var ApiStatisticNumberOfCustomerGet = metagin.Get(
	"getNumberOfCustomer",
	"Number of Customer",
	fmt.Sprintf("/api/%s/motor-garage/statistic/number-of-customer", config.SystemVersion),
	func(ctx metagin.Context[base.Empty, dto.MotorGarageStatistic]) {
		n := relationship.CustomerCaller.CountCustomers(ctx.DB(), ctx.WorkspaceId())

		ctx.OK(&dto.MotorGarageStatistic{
			Label: "number_of_cutomer",
			Value: int(n),
		})
	},
)
