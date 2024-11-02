package main

import (
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage"
)

func main() {
	module_motor_garage.SetupVehicle()
	module_motor_garage.SetupService()
	module_motor_garage.SetupProduct()
	module_motor_garage.SetupDiscount()
	module_motor_garage.SetupInvoice()
	module_motor_garage.SetupInvoiceReport()
	metagin.GenerateTypescript()
	metagin.GenerateOpenAPI()
}
