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

	existServices := make(map[uint]entity.MotorGarageInvoiceService)
	for _, s := range e.Services {
		existServices[s.ID] = s
	}
	for _, s := range r.Services {
		sEntity := entity.MotorGarageInvoiceService{
			MotorGarageService: *s.ToEntity(existServices[s.ID].Service),
		}
		if s.ServiceId != 0 {
			sEntity.ServiceId = &s.ServiceId
		}
		e.Services = append(e.Services, sEntity)
	}

	existProducts := make(map[uint]entity.MotorGarageInvoiceProduct)
	for _, p := range e.Products {
		existProducts[p.ID] = p
	}
	for _, p := range r.Products {
		pEntity := entity.MotorGarageInvoiceProduct{
			MotorGarageProduct: *p.ToEntity(existProducts[p.ID].Product),
			Quantity:           p.Quantity,
		}
		if p.ProductId != 0 {
			pEntity.ProductId = &p.ProductId
		}
		e.Products = append(e.Products, pEntity)
	}

	existDiscounts := make(map[uint]entity.MotorGarageInvoiceDiscount)
	for _, d := range e.Discounts {
		existDiscounts[d.ID] = d
	}
	for _, d := range r.Discounts {
		dEntity := entity.MotorGarageInvoiceDiscount{
			MotorGarageDiscount: *d.ToEntity(existDiscounts[d.ID].Discount),
		}
		if d.DiscountId != 0 {
			dEntity.DiscountId = &d.DiscountId
		}
		e.Discounts = append(e.Discounts, dEntity)
	}
	return e
}
