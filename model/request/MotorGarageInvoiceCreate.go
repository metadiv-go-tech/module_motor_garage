package request

import "github.com/metadiv-go-tech/module_motor_garage/model/entity"

type MotorGarageInvoiceCreate struct {
	Date      int64 `json:"date"`
	VehicleId uint  `json:"vehicle_id"`

	Services  []MotorGarageInvoiceService  `json:"services"`
	Products  []MotorGarageInvoiceProduct  `json:"products"`
	Discounts []MotorGarageInvoiceDiscount `json:"discounts"`
}

func (r *MotorGarageInvoiceCreate) Validate() string {
	if r.Date == 0 {
		return "date is required"
	}
	if r.VehicleId == 0 {
		return "vehicle_id is required"
	}
	return ""
}

func (r *MotorGarageInvoiceCreate) ToEntity(e *entity.MotorGarageInvoice) *entity.MotorGarageInvoice {
	if e == nil {
		e = &entity.MotorGarageInvoice{}
	}
	e.Date = r.Date
	e.VehicleId = r.VehicleId
	for _, s := range r.Services {
		sEntity := entity.MotorGarageInvoiceService{
			MotorGarageService: *s.ToEntity(nil),
		}
		if s.ID != 0 {
			sEntity.ServiceId = &s.ID
		}
		e.Services = append(e.Services, sEntity)
	}
	for _, p := range r.Products {
		pEntity := entity.MotorGarageInvoiceProduct{
			MotorGarageProduct: *p.ToEntity(nil),
			Quantity:           p.Quantity,
		}
		if p.ID != 0 {
			pEntity.ProductId = &p.ID
		}
		e.Products = append(e.Products, pEntity)
	}
	for _, d := range r.Discounts {
		dEntity := entity.MotorGarageInvoiceDiscount{
			MotorGarageDiscount: *d.ToEntity(nil),
		}
		e.Discounts = append(e.Discounts, dEntity)
	}
	return e
}
