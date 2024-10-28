package dto

type MotorGarageService struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Price         uint   `json:"price"`
	PriceAfterTax uint   `json:"price_after_tax"`
}
