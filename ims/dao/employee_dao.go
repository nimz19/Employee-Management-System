package dao

import (
    "database/sql"
    "github.com/nimz19/Employee-Management-System/ims/model" // This import must be used
)

// EmployeeDAO interface
type EmployeeDAO interface {
    CreateEmployee(employee model.Employee) error
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

// Implement at least one method to use the model
func (dao *EmployeeDAOImpl) CreateEmployee(employee model.Employee) error {
    // Example implementation (you can modify this as needed)
    _, err := dao.db.Exec("INSERT INTO employees (first_name, last_name, email, department, salary) VALUES (?, ?, ?, ?, ?)",
        employee.FirstName, employee.LastName, employee.Email, employee.Department, employee.Salary)
    return err
}

// You can add other methods for database operations...
