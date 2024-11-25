package service

import "github.com/metadiv-go-tech/module_motor_garage/model/entity"

var InvoiceService = &invoiceService{}

type invoiceService struct{}

// CalculateInvoiceTotal calculates the total amount of an invoice including services, products, and discounts
func (s *invoiceService) CalculateInvoiceTotal(invoice *entity.MotorGarageInvoice) uint {
	var sum uint = 0
	var discountAmount uint = 0
	var discountPercentage float64 = 1

	// Calculate services total
	if invoice.Services != nil {
		for _, service := range invoice.Services {
			sum += service.PriceAfterTax
		}
	}

	// Calculate products total
	if invoice.Products != nil {
		for _, product := range invoice.Products {
			sum += product.PriceAfterTax * product.Quantity
		}
	}

	// Apply discounts
	if invoice.Discounts != nil {
		for _, discount := range invoice.Discounts {
			discountAmount += discount.DiscountAmount
			discountPercentage *= (1 - float64(discount.DiscountPercentage)/100)
		}
	}

	// Calculate final total after discounts
	return uint(float64(sum-discountAmount) * discountPercentage)
}
