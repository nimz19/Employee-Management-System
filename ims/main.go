package main

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/nimz19/Employee-Management-System/ims/dao" // Correct import path
    "github.com/nimz19/Employee-Management-System/ims/model" // Add this line to import the model package
    _ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

func main() {
    // Connect to the database
    db, err := sql.Open("mysql", "root:StrongP@ssw0rd!@tcp(127.0.0.1:3306)/ims_db")
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }
    defer db.Close() // Ensure the database connection is closed

    // Create a new instance of EmployeeDAO
    employeeDAO := dao.NewEmployeeDAO(db)

    // Use employeeDAO, e.g., to test adding an employee (you might need to create an Employee first)
    emp := model.Employee{
        FirstName:  "Izzah",
        LastName:   "Zafer",
        Email:      "izzahzafer@gmail.com",
        Department: "Software Engineer",
        Salary:     45000.00,
    }

    err = employeeDAO.CreateEmployee(emp) // Call the CreateEmployee method
    if err != nil {
        fmt.Println("Error adding employee:", err)
    } else {
        fmt.Println("Employee added successfully!")
    }

    fmt.Println("Successfully connected to the database!")
}
