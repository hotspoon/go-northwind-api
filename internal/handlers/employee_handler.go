package handlers

import (
	"net/http"
	"northwind-api/internal/models"
	"northwind-api/internal/repositories"
	"northwind-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	Repo *repositories.EmployeeRepository
}

// @Summary Get all employees
// @Description Returns a list of all employees
// @Tags Employees
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Employee
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/employees [get]
func (h *EmployeeHandler) GetAll(c *gin.Context) {
	employees, err := h.Repo.GetAllEmployees(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

// @Summary Get employee by ID
// @Description Returns a single employee by ID
// @Tags Employees
// @Produce json
// @Security BearerAuth
// @Param id path int true "Employee ID"
// @Success 200 {object} models.Employee
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/employees/{id} [get]
func (h *EmployeeHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	employee, err := h.Repo.GetEmployeeByID(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employee)
}

// @Summary Create a new employee
// @Description Creates a new employee
// @Tags Employees
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param employee body models.Employee true "Employee to create"
// @Success 201 {object} models.Employee
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/employees [post]
func (h *EmployeeHandler) Create(c *gin.Context) {
	var emp models.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.Repo.CreateEmployee(c.Request.Context(), emp)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	emp.EmployeeID = int(id)
	c.JSON(http.StatusCreated, gin.H{"message": "employee created successfully"})
}

// @Summary Update an employee
// @Description Updates an existing employee
// @Tags Employees
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Employee ID"
// @Param employee body models.Employee true "Employee to update"
// @Success 200 {object} models.Employee
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/employees/{id} [put]
func (h *EmployeeHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var emp models.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	emp.EmployeeID = utils.ParseInt(id)
	if err := h.Repo.UpdateEmployee(c.Request.Context(), &emp); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "employee updated successfully"})
}

// @Summary Delete an employee
// @Description Deletes an employee by ID
// @Tags Employees
// @Produce json
// @Security BearerAuth
// @Param id path int true "Employee ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/employees/{id} [delete]
func (h *EmployeeHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.Repo.DeleteEmployee(c.Request.Context(), utils.ParseInt(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "employee deleted successfully"})
}
