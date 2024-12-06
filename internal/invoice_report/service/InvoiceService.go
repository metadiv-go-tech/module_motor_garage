package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/metadiv-go-tech/module_motor_garage/internal/invoice_report/template"
	"github.com/metadiv-go-tech/module_motor_garage/model/entity"
)

var InvoiceService = &invoiceService{}

type invoiceService struct{}

func (s *invoiceService) GenerateReport(invoice *entity.MotorGarageInvoice, locale string) string {

	html := template.InvoiceTemplate

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

	var sum uint = 0
	var discountAmount uint = 0
	var discountPercentage float64 = 1

	if invoice.Services != nil {
		sHtml := ""
		total := uint(0)
		for _, service := range invoice.Services {
			serviceDto := service.ToDTO()
			sHtml += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>$%.2f</td><td>$%.2f</td></tr>",
				serviceDto.Name,
				serviceDto.Description,
				float64(service.Price)/100,
				float64(service.PriceAfterTax)/100)
			total += service.PriceAfterTax
			sum += service.PriceAfterTax
		}
		html = strings.Replace(html, "{{service_items}}", sHtml, -1)
		html = strings.ReplaceAll(html, "{{services_total}}", fmt.Sprintf("$%.2f", float64(total)/100))
	} else {
		html = strings.Replace(html, "{{service_items}}", "", -1)
	}

	if invoice.Products != nil {
		pHtml := ""
		total := uint(0)
		for _, product := range invoice.Products {
			productDto := product.ToDTO()
			pHtml += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%d</td><td>$%.2f</td><td>$%.2f</td></tr>",
				productDto.Name, productDto.Description, productDto.Quantity, float64(product.Price)/100, float64(product.PriceAfterTax)/100)
			total += product.PriceAfterTax * productDto.Quantity
			sum += product.PriceAfterTax * productDto.Quantity
		}
		html = strings.Replace(html, "{{product_items}}", pHtml, -1)
		html = strings.ReplaceAll(html, "{{products_total}}", fmt.Sprintf("$%.2f", float64(total)/100))
	} else {
		html = strings.Replace(html, "{{product_items}}", "", -1)
	}

	if invoice.Discounts != nil {
		dHtml := ""
		total := uint(0)
		percentage := float64(1)
		for _, discount := range invoice.Discounts {
			discountDto := discount.ToDTO()
			dHtml += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>$%.2f</td><td>%.2f%%</td></tr>",
				discountDto.Name, discountDto.Description, float64(discount.DiscountAmount)/100, float64(discountDto.DiscountPercentage))
			total += discount.DiscountAmount
			percentage *= (1 - float64(discountDto.DiscountPercentage)/100)
			discountAmount += discount.DiscountAmount
			discountPercentage *= (1 - float64(discountDto.DiscountPercentage)/100)
		}
		html = strings.Replace(html, "{{discount_items}}", dHtml, -1)
		html = strings.Replace(html, "{{discounts_total}}", fmt.Sprintf("$%.2f", float64(total)/100), -1)
		html = strings.ReplaceAll(html, "{{discounts_percentage}}", fmt.Sprintf("%.2f", (1-percentage)*100))
	} else {
		html = strings.Replace(html, "{{discount_items}}", "", -1)
	}

	invoiceTotal := uint(float64(sum-discountAmount) * discountPercentage)
	html = strings.Replace(html, "{{discounts_amount}}", fmt.Sprintf("-$%.2f", (float64(sum)-float64(invoiceTotal))/100), -1)
	html = strings.Replace(html, "{{invoice_total}}", fmt.Sprintf("$%.2f", float64(invoiceTotal)/100), -1)
	return html
}

func (s *invoiceService) GenerateBlankReport() string {

	html := template.InvoiceTemplate

	html = strings.Replace(html, "{{invoice_number}}", "", -1)
	html = strings.Replace(html, "{{invoice_date}}", "", -1)

	html = strings.Replace(html, "{{customer_name}}", "", -1)
	html = strings.Replace(html, "{{vehicle_name}}", "", -1)
	html = strings.Replace(html, "{{vehicle_year}}", "", -1)
	html = strings.Replace(html, "{{vehicle_odometer}}", "", -1)
	html = strings.Replace(html, "{{vehicle_registration}}", "", -1)

	sHtml := ""
	for i := 0; i < 5; i++ {
		sHtml += "<tr><td></td><td></td><td></td><td></td></tr>"
	}
	html = strings.Replace(html, "{{service_items}}", sHtml, -1)
	html = strings.ReplaceAll(html, "{{services_total}}", "$")

	pHtml := ""
	for i := 0; i < 5; i++ {
		pHtml += "<tr><td></td><td></td><td></td><td></td><td></td></tr>"
	}
	html = strings.Replace(html, "{{product_items}}", pHtml, -1)
	html = strings.ReplaceAll(html, "{{products_total}}", "$")

	dHtml := ""
	for i := 0; i < 5; i++ {
		dHtml += "<tr><td></td><td></td><td></td><td></td></tr>"
		html = strings.Replace(html, "{{discount_items}}", dHtml, -1)
	}
	html = strings.ReplaceAll(html, "{{discounts_total}}", "$")
	html = strings.ReplaceAll(html, "{{discounts_percentage}}", "")

	html = strings.Replace(html, "{{discounts_amount}}", "$", -1)
	html = strings.Replace(html, "{{invoice_total}}", "$", -1)
	return html
}
