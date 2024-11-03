package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/metadiv-go-tech/module_motor_garage/internal/test_report/template"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

var ReportService = new(reportService)

type reportService struct{}

func (s *reportService) GenerateReport(invoice *entity.MotorGarageInvoice, locale string) string {

	html := template.TestReportTemplate

	html = strings.Replace(html, "{{invoice_number}}", fmt.Sprintf("#%d", invoice.ID), -1)
	html = strings.Replace(html, "{{invoice_date}}", time.Unix(invoice.Date, 0).Format("02/01/2006"), -1)

	if invoice.Vehicle != nil && invoice.Vehicle.Customer != nil && invoice.Vehicle.Customer.ContactPerson != nil {
		contactPerson := invoice.Vehicle.Customer.ContactPerson.ToDTO()
		html = strings.Replace(html, "{{customer_name}}", contactPerson.FirstName+" "+contactPerson.LastName, -1)
	} else {
		html = strings.Replace(html, "{{customer_name}}", "-", -1)
	}

	if invoice.Vehicle != nil {
		html = strings.Replace(html, "{{vehicle_name}}", invoice.Vehicle.Name, -1)
		html = strings.Replace(html, "{{vehicle_year}}", fmt.Sprintf("%d", invoice.Vehicle.Year), -1)
		html = strings.Replace(html, "{{vehicle_odometer}}", fmt.Sprintf("%d", invoice.Vehicle.Odometer), -1)
		vehicleDto := invoice.Vehicle.ToDTO(locale)
		html = strings.Replace(html, "{{vehicle_registration}}", vehicleDto.Registration, -1)
	} else {
		html = strings.Replace(html, "{{vehicle_name}}", "-", -1)
		html = strings.Replace(html, "{{vehicle_year}}", "-", -1)
		html = strings.Replace(html, "{{vehicle_odometer}}", "-", -1)
	}

	return html
}
