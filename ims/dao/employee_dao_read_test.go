//To test ReadEmployee functionality
//Includes edge cases like non-existing employees
//Expected behavior: The sql.ErrNoRows error is returned by QueryRow().Scan()
//to signal that there was no record found

package dao

import (
	"regexp"
	"testing"
	"fmt"
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nimz19/Employee-Management-System/ims/model"
)

func TestReadEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	employeeDAO := NewEmployeeDAO(db)
	expectedEmployee := model.Employee{
		ID:         1,
		FirstName:  "Izzah",
		LastName:   "Zafer",
		Email:      "izzahzafer@gmail.com",
		Department: "Engineering",
		Salary:     50000,
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, first_name, last_name, email, department, salary FROM employees WHERE id = ?")).
		WithArgs(expectedEmployee.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "department", "salary"}).
			AddRow(expectedEmployee.ID, expectedEmployee.FirstName, expectedEmployee.LastName, expectedEmployee.Email, expectedEmployee.Department, expectedEmployee.Salary))

	emp, err := employeeDAO.ReadEmployee(expectedEmployee.ID)
	if err != nil {
		t.Errorf("Error was not expected: %v", err)
	}

	if emp != expectedEmployee {
		t.Errorf("Expected %+v, got %+v", expectedEmployee, emp)
	}
}

func TestReadNonExistentEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	employeeDAO := NewEmployeeDAO(db)
	nonExistentID := 999 // Use an ID that is not in the mock database

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, first_name, last_name, email, department, salary FROM employees WHERE id = ?")).
		WithArgs(nonExistentID).
		WillReturnError(sql.ErrNoRows)

	_, err = employeeDAO.ReadEmployee(nonExistentID)
	if err == nil || err.Error() != fmt.Sprintf("No employee found with ID %d", nonExistentID) {
		t.Errorf("Expected error for non-existent employee, got: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
