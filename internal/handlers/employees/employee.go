package employees

import (
	. "../../models/employee"
	service "../../services/employees"
	"encoding/json"
	"net/http"
)

func PostEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Employee
	json.NewDecoder(r.Body).Decode(&employee)
	err := service.AddNewEmployee(&employee)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(employee)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	employees, err := service.GetEmployees()
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(employees)
}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	var id uint64
	json.NewDecoder(r.Body).Decode(&id)
	employee, err := service.GetEmployeeById(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(employee)
}

func PutEmployee(w http.ResponseWriter, r *http.Request) {

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	var id uint64
	json.NewDecoder(r.Body).Decode(&id)
	err := service.DeleteEmployee(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
}
