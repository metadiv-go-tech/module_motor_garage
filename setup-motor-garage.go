package module_motor_garage

import (
	"github.com/metadiv-go-tech/metagin/v2"
	bookingHandler "github.com/metadiv-go-tech/module_motor_garage/internal/booking/handler"
	discountHandler "github.com/metadiv-go-tech/module_motor_garage/internal/discount/handler"
	invoiceHandler "github.com/metadiv-go-tech/module_motor_garage/internal/invoice/handler"
	invoiceReportHandler "github.com/metadiv-go-tech/module_motor_garage/internal/invoice_report/handler"
	productHandler "github.com/metadiv-go-tech/module_motor_garage/internal/product/handler"
	serviceHandler "github.com/metadiv-go-tech/module_motor_garage/internal/service/handler"
	technicianHandler "github.com/metadiv-go-tech/module_motor_garage/internal/technician/handler"
	testReportHandler "github.com/metadiv-go-tech/module_motor_garage/internal/test_report/handler"
	vehicleHandler "github.com/metadiv-go-tech/module_motor_garage/internal/vehicle/handler"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

func SetupMotorGarage() {
	metagin.Migrate(
		&entity.MotorGarageBooking{},
		&entity.MotorGarageVehicle{},
		&entity.MotorGarageService{},
		&entity.MotorGarageProduct{},
		&entity.MotorGarageDiscount{},
		&entity.MotorGarageInvoice{},
		&entity.MotorGarageInvoiceService{},
		&entity.MotorGarageInvoiceProduct{},
		&entity.MotorGarageInvoiceDiscount{},
	)

	metagin.RegisterHandler(
		// Vehicle
		vehicleHandler.ApiMotorGarageVehicleCreate,
		vehicleHandler.ApiMotorGarageVehicleList,
		vehicleHandler.ApiMotorGarageVehicleGet,
		vehicleHandler.ApiMotorGarageVehicleUpdate,
		vehicleHandler.ApiMotorGarageVehicleDelete,

		// Service
		serviceHandler.ApiMotorGarageServiceCreate,
		serviceHandler.ApiMotorGarageServiceList,
		serviceHandler.ApiMotorGarageServiceGet,
		serviceHandler.ApiMotorGarageServiceUpdate,
		serviceHandler.ApiMotorGarageServiceDelete,

		// Product
		productHandler.ApiMotorGarageProductCreate,
		productHandler.ApiMotorGarageProductList,
		productHandler.ApiMotorGarageProductGet,
		productHandler.ApiMotorGarageProductUpdate,
		productHandler.ApiMotorGarageProductDelete,

		// Discount
		discountHandler.ApiMotorGarageDiscountCreate,
		discountHandler.ApiMotorGarageDiscountList,
		discountHandler.ApiMotorGarageDiscountGet,
		discountHandler.ApiMotorGarageDiscountUpdate,
		discountHandler.ApiMotorGarageDiscountDelete,

		// Invoice
		invoiceHandler.ApiMotorGarageInvoiceCreate,
		invoiceHandler.ApiMotorGarageInvoiceList,
		invoiceHandler.ApiMotorGarageInvoiceGet,
		invoiceHandler.ApiMotorGarageInvoiceUpdate,
		invoiceHandler.ApiMotorGarageInvoiceDelete,

		// Booking
		bookingHandler.ApiMotorGarageBookingCreate,
		bookingHandler.ApiMotorGarageBookingList,
		bookingHandler.ApiMotorGarageBookingGet,
		bookingHandler.ApiMotorGarageBookingUpdate,
		bookingHandler.ApiMotorGarageBookingDelete,

		// Technician
		technicianHandler.ApiMotorGarageTechnicianCreate,
		technicianHandler.ApiMotorGarageTechnicianList,
		technicianHandler.ApiMotorGarageTechnicianGet,
		technicianHandler.ApiMotorGarageTechnicianUpdate,
		technicianHandler.ApiMotorGarageTechnicianDelete,

		// Invoice Report
		invoiceReportHandler.ApiInvoiceReport,

		// Test Report
		testReportHandler.ApiTestReportPrint,
	)
}
