# Employee-Management-System

This project is a backend system for managing employee information, built with Go. It provides a RESTful API to handle employee data, including adding, updating, deleting, and retrieving employees. The backend connects to a MySQL database, and it’s designed with modularity and testability in mind.

Table of Contents
Features
Tech Stack
Project Structure
Setup and Installation
Environment Variables
API Endpoints
Code Functionality and Flow
Testing
Future Improvements
Features
CRUD Operations: Create, read, update, and delete employees.
Modular Design: Separation of concerns with a structured dao (Data Access Object), service (business logic), and API layers.
Error Handling: Graceful error handling with descriptive messages and logging.
Context-Based Timeout: Uses Go’s context package to handle timeouts on database operations.
Unit Tests: Comprehensive testing for core functionalities.
Tech Stack
Go: Core language for the backend.
MySQL: Relational database for storing employee data.
Gorilla Mux: Router for handling API requests.
Docker: (Optional) For containerized development and deployment.
Project Structure
plaintext
Copy code
ims-backend/
├── dao/                   # Data access layer
│   ├── employee_dao.go    # Employee DAO for database operations
├── model/                 # Data models
│   ├── employee.go        # Employee struct
├── service/               # Business logic layer
│   ├── employee_service.go# Service for employee operations
├── ui/                    # CLI or frontend support if needed
├── main.go                # Entry point
├── go.mod                 # Go module dependencies
└── README.md              # Documentation
Setup and Installation
Clone the Repository:

bash
Copy code
git clone https://github.com/yourusername/ims-backend.git
cd ims-backend
Configure Environment Variables:

Refer to the .env.example file for required environment variables and create a .env file in the root directory.
Run MySQL (if not using Docker):

Ensure you have a MySQL server running with a database named ims_db.
Install Dependencies:

bash
Copy code
go mod tidy
Start the Server:

bash
Copy code
go run main.go
Environment Variables
Variable	Description
DB_USER	Database username
DB_PASSWORD	Database password
DB_NAME	Database name (default: ims_db)
DB_HOST	Database host (default: localhost)
API Endpoints
Endpoint	Method	Description
/employees	POST	Add a new employee
/employees/{id}	GET	Get an employee by ID
/employees/{id}	PUT	Update an employee by ID
/employees/{id}	DELETE	Delete an employee by ID
/employees	GET	List all employees
Code Functionality and Flow
main.go
The main.go file is the entry point for the application. It:

Sets up the database connection.
Configures the API routes using Gorilla Mux.
Initiates the EmployeeService and EmployeeDAO instances.
Starts the server on the specified port.
model/employee.go
Defines the Employee struct that represents the employee data model, with fields like ID, FirstName, LastName, Email, Department, and Salary. This struct matches the database schema.

dao/employee_dao.go
The employee_dao.go file handles all database interactions for the Employee entity:

CreateEmployee: Inserts a new employee into the database.
ReadEmployee: Retrieves an employee by ID.
UpdateEmployee: Updates an employee's information.
DeleteEmployee: Deletes an employee by ID.
GetAllEmployees: Returns a list of all employees.
These methods use SQL queries and return either the expected data or an error if the operation fails.

service/employee_service.go
This file encapsulates business logic, ensuring data is validated before interacting with the DAO. It:

Receives requests from main.go.
Validates data (e.g., checks that Salary is positive, Email is present).
Passes validated data to EmployeeDAO for database interaction.
Returns errors if data validation or DAO operations fail, with contextual error messages.
ui/
(Optional) This folder could house CLI interactions if using a CLI for development testing.

Testing
Unit Tests:

Unit tests are written for the dao and service layers to verify CRUD functionality.
Tests cover edge cases, such as non-existent employee retrieval and handling duplicate emails.
Running Tests:

bash
Copy code
go test ./...
Testing Frameworks:

The standard Go testing package is used, with additional libraries for mocking database operations if needed.
