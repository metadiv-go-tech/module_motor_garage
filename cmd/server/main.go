package main

import (
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/mod_dev_auth/v2"
	"github.com/metadiv-go-tech/module_motor_garage"
	"github.com/metadiv-go-tech/module_relationship"
	"github.com/metadiv-go-tech/module_system_data"
)

func main() {
	mod_dev_auth.QuickSetup()
	module_system_data.SetupCountry()
	module_system_data.SetupAddress()
	module_relationship.SetupCustomer()
	module_motor_garage.SetupVehicle()
	module_motor_garage.SetupService()
	module_motor_garage.SetupProduct()
	module_motor_garage.SetupDiscount()
	module_motor_garage.SetupInvoice()
	module_motor_garage.SetupInvoiceReport()
	module_motor_garage.SetupTestReport()
	metagin.Run()
}
