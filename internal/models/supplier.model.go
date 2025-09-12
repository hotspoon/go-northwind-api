package models

type Supplier struct {
	SupplierID   int64   `json:"supplier_id" db:"SupplierID"`
	CompanyName  string  `json:"company_name" db:"CompanyName"`
	ContactName  *string `json:"contact_name" db:"ContactName"`
	ContactTitle *string `json:"contact_title" db:"ContactTitle"`
	Address      *string `json:"address" db:"Address"`
	City         *string `json:"city" db:"City"`
	Region       *string `json:"region" db:"Region"`
	PostalCode   *string `json:"postal_code" db:"PostalCode"`
	Country      *string `json:"country" db:"Country"`
	Phone        *string `json:"phone" db:"Phone"`
	Fax          *string `json:"fax" db:"Fax"`
	HomePage     *string `json:"homepage" db:"HomePage"`
}
