package model

import (
    "testing"
)

// TestEmployeeStruct initializes an Employee struct and checks field values
func TestEmployeeStruct(t *testing.T) {
    // Directly initialize the Employee struct
    emp := Employee{
        ID:         1,
        FirstName:  "Izzah",
        LastName:   "Zafer",
        Email:      "izzahzafer@gmail.com",
        Department: "Software Engineering",
        Salary:     50000.00,
    }

    // Check each field for expected values
    if emp.ID != 1 {
        t.Errorf("Expected ID 1, got %d", emp.ID)
    }
    if emp.FirstName != "Izzah" {
        t.Errorf("Expected FirstName Izzah, got %s", emp.FirstName)
    }
    if emp.LastName != "Zafer" {
        t.Errorf("Expected LastName Zafer, got %s", emp.LastName)
    }
    if emp.Email != "izzahzafer@gmail.com" {
        t.Errorf("Expected Email izzahzafer@gmail.com, got %s", emp.Email)
    }
    if emp.Department != "Software Engineering" {
        t.Errorf("Expected Department Software Engineering, got %s", emp.Department)
    }
    if emp.Salary != 50000.00 {
        t.Errorf("Expected Salary 50000.00, got %f", emp.Salary)
    }
}

// TestEmployeeString tests the String() method of Employee
func TestEmployeeString(t *testing.T) {
    emp := Employee{
        ID:         1,
        FirstName:  "Izzah",
        LastName:   "Zafer",
        Email:      "izzahzafer@gmail.com",
        Department: "Software Engineering",
        Salary:     50000.00,
    }

    expectedOutput := "ID: 1, Name: Izzah Zafer, Email: izzahzafer@gmail.com, Department: Software Engineering, Salary: 50000.00"
    if emp.String() != expectedOutput {
        t.Errorf("Expected output: %s, got: %s", expectedOutput, emp.String())
    }
}

