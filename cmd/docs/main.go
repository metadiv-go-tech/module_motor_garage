package main

import (
	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/module_motor_garage"
)

func main() {
	module_motor_garage.SetupMotorGarage()
	metagin.GenerateTypescript()
	metagin.GenerateOpenAPI()
}
