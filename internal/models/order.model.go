package models

type Order struct {
	OrderID        int64    `json:"order_id" db:"OrderID"`                // INTEGER, PK, Auto Increment, Not Null
	CustomerID     *string  `json:"customer_id" db:"CustomerID"`          // TEXT, nullable
	EmployeeID     *int64   `json:"employee_id" db:"EmployeeID"`          // INTEGER, nullable
	OrderDate      *string  `json:"order_date" db:"OrderDate"`            // DATETIME, nullable (use *time.Time if you want time type)
	RequiredDate   *string  `json:"required_date" db:"RequiredDate"`      // DATETIME, nullable
	ShippedDate    *string  `json:"shipped_date" db:"ShippedDate"`        // DATETIME, nullable
	ShipVia        *int64   `json:"ship_via" db:"ShipVia"`                // INTEGER, nullable
	Freight        *float64 `json:"freight" db:"Freight"`                 // NUMERIC, nullable (default 0)
	ShipName       *string  `json:"ship_name" db:"ShipName"`              // TEXT, nullable
	ShipAddress    *string  `json:"ship_address" db:"ShipAddress"`        // TEXT, nullable
	ShipCity       *string  `json:"ship_city" db:"ShipCity"`              // TEXT, nullable
	ShipRegion     *string  `json:"ship_region" db:"ShipRegion"`          // TEXT, nullable
	ShipPostalCode *string  `json:"ship_postal_code" db:"ShipPostalCode"` // TEXT, nullable
	ShipCountry    *string  `json:"ship_country" db:"ShipCountry"`        // TEXT, nullable
}

type OrderDetail struct {
	OrderID   int64   `json:"order_id" db:"OrderID"`     // INTEGER, Not Null
	ProductID int64   `json:"product_id" db:"ProductID"` // INTEGER, Not Null
	UnitPrice float64 `json:"unit_price" db:"UnitPrice"` // NUMERIC, Not Null (default 0)
	Quantity  int     `json:"quantity" db:"Quantity"`    // INTEGER, Not Null (default 1)
	Discount  float32 `json:"discount" db:"Discount"`    // REAL, Not Null (default 0)
}
