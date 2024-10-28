package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/nimz19/Employee-Management-System/ims/dao"
	"github.com/nimz19/Employee-Management-System/ims/model"
	"github.com/nimz19/Employee-Management-System/ims/service"
)

func main() {
	// Set up context with timeout for establishing database connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Set up database connection
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ims_db")
	if err != nil {
		log.Fatalf("Error creating database connection: %v", err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Database connection successfully established.")

	// Set up DAO and Service
	employeeDAO := dao.NewEmployeeDAO(db)
	employeeService := service.NewEmployeeService(employeeDAO)

	// Set up router and API routes
	router := mux.NewRouter()
	router.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		addEmployeeHandler(w, r, employeeService)
	}).Methods("POST")

	router.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		getAllEmployeesHandler(w, r, employeeService)
	}).Methods("GET")

	router.HandleFunc("/employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		getEmployeeByIDHandler(w, r, employeeService)
	}).Methods("GET")

	// Serve static files for the frontend using absolute path
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("C:/Users/User/Documents/Employee-Management-System/frontend"))))

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Handler to add a new employee
func addEmployeeHandler(w http.ResponseWriter, r *http.Request, service *service.EmployeeServiceImpl) {
	var emp model.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := service.CreateEmployee(ctx, emp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee added successfully"})
}

// Handler to retrieve all employees
func getAllEmployeesHandler(w http.ResponseWriter, r *http.Request, service *service.EmployeeServiceImpl) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	employees, err := service.GetAllEmployees(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

// Handler to get an employee by ID
func getEmployeeByIDHandler(w http.ResponseWriter, r *http.Request, service *service.EmployeeServiceImpl) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	emp, err := service.GetEmployeeByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}
