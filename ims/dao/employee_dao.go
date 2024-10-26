package dao

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nimz19/Employee-Management-System/ims/model" // This import must be used
)

// EmployeeDAO interface
type EmployeeDAO interface {
	CreateEmployee(ctx context.Context, employee model.Employee) error
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

// Implement ReadEmployee
func (dao *EmployeeDAOImpl) ReadEmployee(id int) (model.Employee, error) {
	var employee model.Employee
	query := "SELECT id, first_name, last_name, email, department, salary FROM employees WHERE id = ?"

	row := dao.db.QueryRow(query, id)
	err := row.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Email, &employee.Department, &employee.Salary)
	if err != nil {
		if err == sql.ErrNoRows {
			return employee, fmt.Errorf("No employee found with ID %d", id)
		}
		return employee, fmt.Errorf("Error retrieving employee: %w", err)
	}
	return employee, nil
}

// Implement UpdateEmployee
func (dao *EmployeeDAOImpl) UpdateEmployee(employee model.Employee) error {
	query := "UPDATE employees SET first_name = ?, last_name = ?, email = ?, department = ?, salary = ? WHERE id = ?"
	_, err := dao.db.Exec(query, employee.FirstName, employee.LastName, employee.Email, employee.Department, employee.Salary, employee.ID)
	if err != nil {
		return fmt.Errorf("Error updating employee: %w", err)
	}
	return nil
}

// Implement DeleteEmployee
func (dao *EmployeeDAOImpl) DeleteEmployee(id int) error {
	query := "DELETE FROM employees WHERE id = ?"
	_, err := dao.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error deleting employee: %w", err)
	}
	return nil
}

// Implement GetAllEmployees
func (dao *EmployeeDAOImpl) GetAllEmployees() ([]model.Employee, error) {
	query := "SELECT id, first_name, last_name, email, department, salary FROM employees"
	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving employees: %w", err)
	}
	defer rows.Close()

	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		if err := rows.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Email, &employee.Department, &employee.Salary); err != nil {
			return nil, fmt.Errorf("Error scanning employee: %w", err)
		}
		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over employees: %w", err)
	}
	return employees, nil
}
