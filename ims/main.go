package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "time"

    "github.com/nimz19/Employee-Management-System/ims/dao"      // Correct import path for DAO
    "github.com/nimz19/Employee-Management-System/ims/model"    // Import the model package
    "github.com/nimz19/Employee-Management-System/ims/service"  // Import the service package
    _ "github.com/go-sql-driver/mysql"                          // Import the MySQL driver
)

func main() {
    // Connect to the database
    db, err := sql.Open("mysql", "root:StrongP@ssw0rd!@tcp(127.0.0.1:3306)/ims_db")
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }
    defer db.Close() // Ensure the database connection is closed

    // Ping the database to verify connection
    if err := db.Ping(); err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

    // Create a new instance of EmployeeDAO
    employeeDAO := dao.NewEmployeeDAO(db)

    // Create a new instance of EmployeeService with EmployeeDAO
    employeeService := service.NewEmployeeService(employeeDAO)

    // Define an employee to add, demonstrating business logic validations
    emp := model.Employee{
        FirstName:  "Vicky",
        LastName:   "Crudzs",
        Email:      "vickycrudz@gmail.com",
        Department: "HR",
        Salary:     30000.00,
    }

    // Use context with timeout for the CreateEmployee method
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel() // Ensure the cancel function is called to avoid context leaks

    // Call the EmployeeService to create the employee, allowing for validation and business logic
    err = employeeService.CreateEmployee(ctx, emp)
    if err != nil {
        fmt.Println("Error adding employee:", err)
    } else {
        fmt.Println("Employee added successfully!")
    }

    // Test the ReadEmployee method through EmployeeService
    // Replace this ID with the actual ID of the employee you just added if needed
    employeeID := 1 // Use a valid ID or dynamically retrieve it

    // Use context with timeout for the ReadEmployee method
    ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel() // Ensure the cancel function is called to avoid context leaks

    // Call the EmployeeService to read an employee by ID
    employee, err := employeeService.GetEmployeeByID(ctx, employeeID)
    if err != nil {
        fmt.Println("Error reading employee:", err)
    } else {
        fmt.Printf("Employee found: %+v\n", employee) // Print employee details
    }

    fmt.Println("Successfully connected to the database!")
}
