const API_URL = "http://localhost:8080"; // Adjust this to your Go backend’s URL

// Add employee on form submission
document.getElementById("addEmployeeForm").addEventListener("submit", async (e) => {
    e.preventDefault();

    const employeeData = {
        firstName: document.getElementById("firstName").value,
        lastName: document.getElementById("lastName").value,
        email: document.getElementById("email").value,
        department: document.getElementById("department").value,
        salary: parseFloat(document.getElementById("salary").value)
    };

    const response = await fetch(`${API_URL}/employees`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(employeeData),
    });

    const result = await response.json();
    if (response.ok) {
        alert("Employee added successfully!");
        document.getElementById("addEmployeeForm").reset();
    } else {
        alert(`Failed to add employee: ${result.message}`);
    }
});

// Fetch and display all employees
document.getElementById("fetchEmployees").addEventListener("click", async () => {
    console.log("Fetch Employees button clicked"); // Should log when button is clicked
    const response = await fetch(`${API_URL}/employees`);
    const employees = await response.json();
    console.log(employees); // Should log the employee data

    let output = "<ul>";
    employees.forEach(emp => {
        output += `<li>${emp.FirstName} ${emp.LastName} - ${emp.Department} - $${emp.Salary}</li>`;
    });
    output += "</ul>";

    document.getElementById("employeeOutput").innerHTML = output;
});
