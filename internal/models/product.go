package models

type Product struct {
	ProductID       int     `json:"product_id"`
	ProductName     string  `json:"product_name"`
	SupplierID      int     `json:"supplier_id"`
	CategoryID      int     `json:"category_id"`
	QuantityPerUnit string  `json:"quantity_per_unit"`
	UnitPrice       float64 `json:"unit_price"`
	UnitsInStock    int     `json:"units_in_stock"`
	UnitsOnOrder    int     `json:"units_on_order"`
	ReorderLevel    int     `json:"reorder_level"`
	Discontinued    bool    `json:"discontinued"`
}
