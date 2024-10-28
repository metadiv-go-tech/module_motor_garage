package main

import (
	"github.com/metadiv-go-tech/metagin"
	"github.com/metadiv-go-tech/mod_dev_auth"
	"github.com/metadiv-go-tech/module_motor_garage"
)

func main() {
	mod_dev_auth.QuickSetup()
	module_motor_garage.SetupVehicle()
	module_motor_garage.SetupService()
	module_motor_garage.SetupProduct()
	module_motor_garage.SetupDiscount()
	module_motor_garage.SetupInvoice()
	metagin.Engine.Run()
}
