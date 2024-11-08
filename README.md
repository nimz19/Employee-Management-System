# Employee-Management-System

A command-line Employee Management System backend built in Go. This system enables users to manage employees, supporting operations such as adding new employees, updating records, retrieving details, and listing all employees.

## Features

- **Add New Employees**: Add new employees with details like name, department, email, and salary.
- **Retrieve Employee Details**: View specific employee information by ID.
- **Update Employee Records**: Update an existing employee's information.
- **Delete Employee Records**: Remove an employee from the system (database) by their ID.
- **List All Employees**: Display information about all employees in the database.

## Technology Used

- **Go**: Programming language used.
- **MySQL**: Database used to store employee records.
- **Context**: Used to manage request timeouts and cancellations.
- **_test.go**: For unit testing of core functionalities.

## Project Structure

```plaintext
Employee-Manangement-System/
├── ims-backend/
│   ├── dao/
│   │   └── employee_dao.go
│   ├── model/
│   │   └── employee.go
│   ├── service/
│   │   └── employee_service.go
│   ├── handlers/                # New folder for request handlers
│   │   └── employee_handlers.go # Handler for employee-related HTTP requests
│   ├── ui/
│   │   └── cli.go
│   ├── main.go
│   ├── go.mod
│   └── README.md
└── frontend/
    ├── index.html
    ├── script.js
    └── styles.css

```
## Code Walkthrough

## Code Walkthrough

### `main.go` (Main Program)

- Sets up database connection and initializes services.
- Calls `EmployeeService` for business logic operations and `EmployeeDAO` for data access.

### `employee_dao.go` (Data Access Layer)

- Defines methods to interact with the database using CRUD operations.
- Contains SQL queries for creating, reading, updating, and deleting employee records.
- Handles database connection errors and logs error details for troubleshooting.

### `employee_service.go` (Business Logic Layer)

- Validates employee data before passing it to `EmployeeDAO` for database operations.
- Includes business rules and constraints to ensure the salary is positive and the email is non-empty.
- Provides core methods for managing employees: `CreateEmployee`, `GetEmployeeByID`, `UpdateEmployee`, `DeleteEmployee`, and `GetAllEmployees`.

### `cli.go` (Command Line Interface)

- A command-line text-based menu to interact with the system.

      
 ### Useful Commands
 in command prompt to kill process
 Find port: netstat -aon | findstr :8080
 Kill (find PID number): taskkill /PID 19776 /F 


