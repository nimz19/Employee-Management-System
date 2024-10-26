package ui

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/nimz19/Employee-Management-System/ims/model"
	"github.com/nimz19/Employee-Management-System/ims/service"
)

// StartCLI starts the CLI for interacting with the Employee Management System
func StartCLI(employeeService service.EmployeeService) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Display menu options
		fmt.Println("\nEmployee Management System")
		fmt.Println("1. Add Employee")
		fmt.Println("2. View Employee by ID")
		fmt.Println("3. Update Employee")
		fmt.Println("4. Delete Employee by ID")
		fmt.Println("5. List All Employees")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addEmployee(scanner, employeeService)
		case "2":
			viewEmployee(scanner, employeeService)
		case "3":
			updateEmployee(scanner, employeeService)
		case "4":
			deleteEmployee(scanner, employeeService)
		case "5":
			listAllEmployees(employeeService)
		case "6":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// Helper function to add an employee
func addEmployee(scanner *bufio.Scanner, employeeService service.EmployeeService) {
	fmt.Println("Enter employee details:")
	fmt.Print("First Name: ")
	scanner.Scan()
	firstName := scanner.Text()

	fmt.Print("Last Name: ")
	scanner.Scan()
	lastName := scanner.Text()

	fmt.Print("Email: ")
	scanner.Scan()
	email := scanner.Text()

	fmt.Print("Department: ")
	scanner.Scan()
	department := scanner.Text()

	fmt.Print("Salary: ")
	scanner.Scan()
	salaryStr := scanner.Text()
	salary, err := strconv.ParseFloat(salaryStr, 64)
	if err != nil {
		fmt.Println("Invalid salary. Please enter a numeric value.")
		return
	}

	emp := model.Employee{
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		Department: department,
		Salary:     salary,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := employeeService.CreateEmployee(ctx, emp); err != nil {
		fmt.Println("Error adding employee:", err)
	} else {
		fmt.Println("Employee added successfully!")
	}
}

// Helper function to view an employee by ID
func viewEmployee(scanner *bufio.Scanner, employeeService service.EmployeeService) {
	fmt.Print("Enter Employee ID: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a numeric value.")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	employee, err := employeeService.GetEmployeeByID(ctx, id)
	if err != nil {
		fmt.Println("Error fetching employee:", err)
	} else {
		fmt.Printf("Employee Details: %+v\n", employee)
	}
}

// Helper function to update an employee
func updateEmployee(scanner *bufio.Scanner, employeeService service.EmployeeService) {
	fmt.Print("Enter Employee ID to update: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a numeric value.")
		return
	}

	fmt.Print("First Name: ")
	scanner.Scan()
	firstName := scanner.Text()

	fmt.Print("Last Name: ")
	scanner.Scan()
	lastName := scanner.Text()

	fmt.Print("Email: ")
	scanner.Scan()
	email := scanner.Text()

	fmt.Print("Department: ")
	scanner.Scan()
	department := scanner.Text()

	fmt.Print("Salary: ")
	scanner.Scan()
	salaryStr := scanner.Text()
	salary, err := strconv.ParseFloat(salaryStr, 64)
	if err != nil {
		fmt.Println("Invalid salary. Please enter a numeric value.")
		return
	}

	emp := model.Employee{
		ID:         id,
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		Department: department,
		Salary:     salary,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := employeeService.UpdateEmployee(ctx, emp); err != nil {
		fmt.Println("Error updating employee:", err)
	} else {
		fmt.Println("Employee updated successfully!")
	}
}

// Helper function to delete an employee by ID
func deleteEmployee(scanner *bufio.Scanner, employeeService service.EmployeeService) {
	fmt.Print("Enter Employee ID to delete: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a numeric value.")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := employeeService.DeleteEmployee(ctx, id); err != nil {
		fmt.Println("Error deleting employee:", err)
	} else {
		fmt.Println("Employee deleted successfully!")
	}
}

// Helper function to list all employees
func listAllEmployees(employeeService service.EmployeeService) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	employees, err := employeeService.GetAllEmployees(ctx)
	if err != nil {
		fmt.Println("Error fetching employees:", err)
	} else {
		fmt.Println("Employee List:")
		for _, emp := range employees {
			fmt.Printf("%+v\n", emp)
		}
	}
}
