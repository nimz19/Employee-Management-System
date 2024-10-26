//To separate the business logic from the data access layer (DAO)
//Benefial for
//1. Encapsulation of Business Logic
//2. Modular Code
//3. Easier testing

package service

import (
	"context"
	"fmt"
	"github.com/nimz19/Employee-Management-System/ims/dao"
	"github.com/nimz19/Employee-Management-System/ims/model"
)

// EmployeeService defines methods for employee operations with business logic
type EmployeeService interface {
	CreateEmployee(ctx context.Context, employee model.Employee) error
	GetEmployeeByID(ctx context.Context, id int) (model.Employee, error)
	UpdateEmployee(ctx context.Context, employee model.Employee) error
	DeleteEmployee(ctx context.Context, id int) error
	GetAllEmployees(ctx context.Context) ([]model.Employee, error)
}

// EmployeeServiceImpl is the implementation of EmployeeService
type EmployeeServiceImpl struct {
	employeeDAO dao.EmployeeDAO
}

// NewEmployeeService creates a new EmployeeService
func NewEmployeeService(employeeDAO dao.EmployeeDAO) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{employeeDAO: employeeDAO}
}

// CreateEmployee validates employee data and adds them to the database
func (s *EmployeeServiceImpl) CreateEmployee(ctx context.Context, employee model.Employee) error {
	// Example business logic: validate employee data
	if employee.Salary <= 0 {
		return fmt.Errorf("salary must be positive")
	}
	if employee.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}

	// Call the DAO layer to create the employee
	return s.employeeDAO.CreateEmployee(ctx, employee)
}

// GetEmployeeByID fetches an employee by their ID
func (s *EmployeeServiceImpl) GetEmployeeByID(ctx context.Context, id int) (model.Employee, error) {
	// Example validation: ensure ID is valid
	if id <= 0 {
		return model.Employee{}, fmt.Errorf("invalid employee ID")
	}

	// Call the DAO layer to get the employee by ID
	return s.employeeDAO.ReadEmployee(id)
}

// UpdateEmployee updates an existing employee in the database
func (s *EmployeeServiceImpl) UpdateEmployee(ctx context.Context, employee model.Employee) error {
	// Example business logic: ensure ID is set for updating
	if employee.ID <= 0 {
		return fmt.Errorf("employee ID must be set for update")
	}

	// Call the DAO layer to update the employee
	return s.employeeDAO.UpdateEmployee(employee)
}

// DeleteEmployee deletes an employee by their ID
func (s *EmployeeServiceImpl) DeleteEmployee(ctx context.Context, id int) error {
	// Example validation: ensure ID is valid
	if id <= 0 {
		return fmt.Errorf("invalid employee ID")
	}

	// Call the DAO layer to delete the employee by ID
	return s.employeeDAO.DeleteEmployee(id)
}

// GetAllEmployees retrieves all employees from the database
func (s *EmployeeServiceImpl) GetAllEmployees(ctx context.Context) ([]model.Employee, error) {
	// Call the DAO layer to get all employees
	return s.employeeDAO.GetAllEmployees()
}
