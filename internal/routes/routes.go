// internal/routes/router.go
package routes

import (
	"database/sql"

	"northwind-api/internal/handlers"
	"northwind-api/internal/repositories"
	"northwind-api/internal/utils"

	"github.com/gin-gonic/gin"
)

// Hanya expose yang dibutuhkan layer routes
type ConfigView interface {
	Env() string    // "production" | "staging" | "development"
	APIVer() string // e.g. "v1"
}

type Deps struct {
	DB     *sql.DB
	Config ConfigView
}

func Register(e *gin.Engine, d Deps) {
	// Build shared repos/handlers here (or inside each sub-registrar)
	customerRepo := &repositories.CustomerRepository{DB: d.DB}
	customerHandler := &handlers.CustomerHandler{Repo: customerRepo}

	employeeRepo := &repositories.EmployeeRepository{DB: d.DB}
	employeeHandler := &handlers.EmployeeHandler{Repo: employeeRepo}

	// Swagger (only non-prod)
	RegisterSwagger(e, d.Config)

	// Versioned API group
	api := e.Group("/api/" + d.Config.APIVer())

	// Protected toggle
	var protected *gin.RouterGroup
	if d.Config.Env() == "production" {
		protected = api.Group("", utils.AuthMiddlewareJWT())
	} else {
		protected = api.Group("")
	}

	// Feature groups
	RegisterCustomerRoutes(protected, customerHandler)
	RegisterEmployeeRoutes(protected, employeeHandler)
}
