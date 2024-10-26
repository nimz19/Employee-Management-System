package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "time"

    "github.com/nimz19/Employee-Management-System/ims/dao"  // Correct import path
    "github.com/nimz19/Employee-Management-System/ims/model" // Import the model package
    _ "github.com/go-sql-driver/mysql"                       // Import the MySQL driver
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

    // Use employeeDAO to add a new employee
    emp := model.Employee{
        FirstName:  "Molly",
        LastName:   "Richards",
        Email:      "mollyrichards@gmail.com",
        Department: "Software Engineer",
        Salary:     45000.00,
    }

    // Use context with timeout for the CreateEmployee method
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel() // Ensure the cancel function is called to avoid context leaks

    err = employeeDAO.CreateEmployee(ctx, emp) // Call the CreateEmployee method with context
    if err != nil {
        fmt.Println("Error adding employee:", err)
    } else {
        fmt.Println("Employee added successfully!")
    }

    fmt.Println("Successfully connected to the database!")
}
