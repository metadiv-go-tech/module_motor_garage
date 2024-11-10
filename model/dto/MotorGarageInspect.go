package dto

type MotorGarageInspect struct {
	ID uint `json:"id"`

	InvoiceId uint                `json:"invoice_id"`
	Invoice   *MotorGarageInvoice `json:"invoice"`

	PassItems []string          `json:"pass_items"`
	FailItems []string          `json:"fail_items"`
	KeyValues map[string]string `json:"key_values"`
}
