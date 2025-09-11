package models

type Shipper struct {
	ShipperID   int    `json:"shipper_id,omitempty"`
	CompanyName string `json:"company_name"`
	Phone       string `json:"phone"`
}
