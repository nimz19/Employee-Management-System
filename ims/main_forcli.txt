package main

import (
    "context"
    "database/sql"
    "log"
    "time"

    "github.com/nimz19/Employee-Management-System/ims/dao"
    "github.com/nimz19/Employee-Management-System/ims/service"
    "github.com/nimz19/Employee-Management-System/ims/ui"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Set up context with timeout for establishing database connection
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel() // Ensure the cancel function is called to avoid context leaks

    // Set up database connection with a DSN
    db, err := sql.Open("mysql", "root:StrongP@ssw0rd!@tcp(127.0.0.1:3306)/ims_db")
    if err != nil {
        log.Fatalf("Error creating database connection: %v", err)
    }
    defer func() {
        if err := db.Close(); err != nil {
            log.Fatalf("Error closing database connection: %v", err)
        }
    }()

    // Ping the database to verify connection
    if err := db.PingContext(ctx); err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    log.Println("Database connection successfully established.")

    // Set up DAO and Service with safe context handling
    employeeDAO := dao.NewEmployeeDAO(db)
    employeeService := service.NewEmployeeService(employeeDAO)

    // Start the CLI with EmployeeService
    ui.StartCLI(employeeService)
}
