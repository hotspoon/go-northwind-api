package models

type Product struct {
	ProductID       int     `json:"product_id"`
	ProductName     string  `json:"product_name"`
	SupplierID      *int    `json:"supplier_id,omitempty"`
	CategoryID      *int    `json:"category_id,omitempty"`
	QuantityPerUnit *string `json:"quantity_per_unit,omitempty"`
	UnitPrice       float64 `json:"unit_price"`
	UnitsInStock    int     `json:"units_in_stock"`
	UnitsOnOrder    int     `json:"units_on_order"`
	ReorderLevel    int     `json:"reorder_level"`
	Discontinued    string  `json:"discontinued"`
}

// GetSupplierByProductID mengembalikan SupplierID & CompanyName (minimalis untuk endpoint /products/{id}/supplier)
type ProductSupplier struct {
	SupplierID  int    `json:"supplier_id"`
	CompanyName string `json:"company_name"`
}

// GetCategoryByProductID mengembalikan CategoryID & CategoryName (untuk endpoint /products/{id}/category)
type ProductCategory struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
