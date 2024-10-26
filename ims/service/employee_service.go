//To separate the business logic from the data access layer (DAO)
//Beneficial for
//1. Encapsulation of Business Logic
//2. Modular Code
//3. Easier testing

//++ Added extra validation and error wrapping
//Also added logging for key event and errors
//to track event flow and easier for debugging
//check for specific errors (missing records)

package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

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
	log.Println("Starting CreateEmployee operation")

	// Validate employee data
	if employee.Salary <= 0 {
		log.Printf("Invalid salary: %v", employee.Salary)
		return fmt.Errorf("salary must be positive")
	}
	if employee.Email == "" {
		log.Println("Email cannot be empty")
		return fmt.Errorf("email cannot be empty")
	}

	// Attempt to create the employee
	err := s.employeeDAO.CreateEmployee(ctx, employee)
	if err != nil {
		log.Printf("Error adding employee to DAO layer: %v", err)
		// Wrap the error with additional context
		return fmt.Errorf("failed to add employee: %w", err)
	}

	log.Println("Successfully added employee")
	return nil
}

// GetEmployeeByID fetches an employee by their ID
func (s *EmployeeServiceImpl) GetEmployeeByID(ctx context.Context, id int) (model.Employee, error) {
	log.Printf("Fetching employee with ID: %d", id)

	// Validate ID
	if id <= 0 {
		log.Printf("Invalid employee ID: %d", id)
		return model.Employee{}, fmt.Errorf("invalid employee ID")
	}

	// Retrieve employee from DAO
	employee, err := s.employeeDAO.ReadEmployee(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Employee with ID %d not found", id)
			return model.Employee{}, fmt.Errorf("employee with ID %d not found", id)
		}
		log.Printf("Error retrieving employee from DAO layer: %v", err)
		return model.Employee{}, fmt.Errorf("failed to retrieve employee: %w", err)
	}

	log.Printf("Successfully retrieved employee: %+v", employee)
	return employee, nil
}

// UpdateEmployee updates an existing employee in the database
func (s *EmployeeServiceImpl) UpdateEmployee(ctx context.Context, employee model.Employee) error {
	log.Printf("Updating employee with ID: %d", employee.ID)

	// Validate employee ID
	if employee.ID <= 0 {
		log.Println("Employee ID must be set for update")
		return fmt.Errorf("employee ID must be set for update")
	}

	// Attempt to update the employee
	err := s.employeeDAO.UpdateEmployee(employee)
	if err != nil {
		log.Printf("Error updating employee in DAO layer: %v", err)
		return fmt.Errorf("failed to update employee: %w", err)
	}

	log.Println("Successfully updated employee")
	return nil
}

// DeleteEmployee deletes an employee by their ID
func (s *EmployeeServiceImpl) DeleteEmployee(ctx context.Context, id int) error {
	log.Printf("Deleting employee with ID: %d", id)

	// Validate ID
	if id <= 0 {
		log.Printf("Invalid employee ID: %d", id)
		return fmt.Errorf("invalid employee ID")
	}

	// Attempt to delete the employee
	err := s.employeeDAO.DeleteEmployee(id)
	if err != nil {
		log.Printf("Error deleting employee from DAO layer: %v", err)
		return fmt.Errorf("failed to delete employee: %w", err)
	}

	log.Println("Successfully deleted employee")
	return nil
}

// GetAllEmployees retrieves all employees from the database
func (s *EmployeeServiceImpl) GetAllEmployees(ctx context.Context) ([]model.Employee, error) {
	log.Println("Fetching all employees")

	// Retrieve all employees
	employees, err := s.employeeDAO.GetAllEmployees()
	if err != nil {
		log.Printf("Error retrieving employees from DAO layer: %v", err)
		return nil, fmt.Errorf("failed to retrieve employees: %w", err)
	}

	log.Printf("Successfully retrieved %d employees", len(employees))
	return employees, nil
}
