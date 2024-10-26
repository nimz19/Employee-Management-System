package dao

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nimz19/Employee-Management-System/ims/model" // This import must be used
)

// EmployeeDAO interface
type EmployeeDAO interface {
	CreateEmployee(ctx context.Context, employee model.Employee) error // Updated to include context
	ReadEmployee(id int) (model.Employee, error)
	UpdateEmployee(employee model.Employee) error
	DeleteEmployee(id int) error
	GetAllEmployees() ([]model.Employee, error)
}

// EmployeeDAOImpl struct
type EmployeeDAOImpl struct {
	db *sql.DB
}

// Create a new instance of EmployeeDAOImpl
func NewEmployeeDAO(db *sql.DB) *EmployeeDAOImpl {
	return &EmployeeDAOImpl{db: db}
}

// Implement CreateEmployee with context
func (dao *EmployeeDAOImpl) CreateEmployee(ctx context.Context, employee model.Employee) error {
	query := "INSERT INTO employees (first_name, last_name, email, department, salary) VALUES (?, ?, ?, ?, ?)"
	
	// Execute the query with context
	_, err := dao.db.ExecContext(ctx, query, employee.FirstName, employee.LastName, employee.Email, employee.Department, employee.Salary)
	if err != nil {
		return fmt.Errorf("Error adding employee: %w", err)
	}
	return nil
}

// You can add other methods for database operations...
