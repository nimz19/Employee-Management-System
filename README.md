# Employee-Management-System

A command-line Employee Management System backend built in Go. This system enables users to manage employess, supporting operations such as adding new employees, updating records, retrieving details, and listing all employees.

## Features
-**Add New Employees**: Add new employees with details; name, department, email salary
-**Retrieve Employee Details**: View specific employee information by ID
-**Update Employee Records**: Update an existing employee's information
-**Delete Employee Records**: Remove an employee from the system (database) by their ID
-**List All Employees**: Display information about all employees in the database

## Technology Used
-**Go**: Programming language used
-**MySQL**: Databased used to store employee records
-**Context**: Used to integrate request timeouts and cancellations management
-**_test.go**: For unit testing of core functionalities

## IMS Backend Project Structure

```plaintext
ims-backend/
├── dao/                   # Data access layer for database interactions
│   ├── employee_dao.go    # Employee DAO for CRUD operations on employee data
├── model/                 # Data models representing database entities
│   ├── employee.go        # Employee struct representing an employee record
├── service/               # Business logic layer
│   ├── employee_service.go # Service with validation and business rules for employee operations
├── ui/                    # CLI interface for user interactions
│   ├── cli.go             # CLI for interacting with the employee management system
├── main.go                # Entry point of the application
├── go.mod                 # Go module dependencies
└── README.md              # Project documentation

### Code Walkthrough

