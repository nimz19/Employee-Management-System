package model

import "fmt"

// Employee struct matching the database structure
type Employee struct {
    ID         int
    FirstName  string
    LastName   string
    Email      string
    Department string
    Salary     float64
}

// String method to print employee details
func (e Employee) String() string {
    return fmt.Sprintf("ID: %d, Name: %s %s, Email: %s, Department: %s, Salary: %.2f",
        e.ID, e.FirstName, e.LastName, e.Email, e.Department, e.Salary)
}

//Getters and Setters are not commonly used in go
//Also don't need a constructor