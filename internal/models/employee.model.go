package models

type Employee struct {
	EmployeeID      int    `json:"employee_id"`
	LastName        string `json:"last_name"`
	FirstName       string `json:"first_name"`
	Title           string `json:"title"`
	TitleOfCourtesy string `json:"title_of_courtesy"`
	BirthDate       string `json:"birth_date"`
	HireDate        string `json:"hire_date"`
	Address         string `json:"address"`
	City            string `json:"city"`
	Region          string `json:"region"`
	PostalCode      string `json:"postal_code"`
	Country         string `json:"country"`
	HomePhone       string `json:"home_phone"`
	Extension       string `json:"extension"`
	Photo           []byte `json:"photo"`
	Notes           string `json:"notes"`
	ReportsTo       *int   `json:"reports_to"`
	PhotoPath       string `json:"photo_path"`
}
