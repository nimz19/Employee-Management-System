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

```
## Code Walkthrough

-**main.go** (Main Program):
  - Set up database connection and initializes services
  - Calls EmployeeService for business logic operations and EmployeeDAO for data access
  - 
-**employee_dao.go** (Data Access Layer):
    - Defines methods to interact with the database using CRUD operations
    - States the SQL queries for operations like creating, reading, updating, and deleting     
      employee records
    - Handles database connection errors and loggic error details for troubleshooting
      
-**employee_service.go** (Business Logic Layer):
    - Validates employee data before passing it to EmployeeDAO for database operations
    - includes business rules and constraints, to ensure salary is positive and email is non-  
      empty
    - Provides core methods for managing employees: CreateEmployee, GetEmployeeByID, 
      UpdateEmployee, DeleteEmployee, and GetAllEmployees.
      
-**cli.go** (Command Line Interface):
    -a command line text-based menu to interact with the system      
      
    


