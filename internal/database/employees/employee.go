package employees

import (
	. "../"
	. "../../models/employee"
	"log"
)

var db, _ = OpenConnection()

func AddNewEmployee(employee *Employee) error {
	query := "INSERT INTO employees(position, first_name, last_name, patronymic,login, password, storage_id) VALUES(?,?,?,?,?,?,?)"

	result, err := db.Exec(query,
		employee.ID,
		employee.FirstName,
		employee.LastName,
		employee.Patronymic,
		employee.User.UserName,
		employee.User.Password,
		employee.Storage.ID)

	if err != nil {
		log.Fatal("Can't add user in DB")
	}
	id, _ := result.LastInsertId()
	employee.ID = uint64(id)
}

func GetEmployeeById(id uint64) (Employee, error) {
	query := "SELECT id, first_name, last_name, patronymic, login, storages FROM employees WHERE id=?"

	var employee Employee
	err := db.QueryRow(query, id).Scan(
		&employee.ID,
		&employee.FirstName,
		&employee.LastName,
		&employee.Patronymic,
		&employee.User.UserName,
		&employee.Storage.ID)

	if err != nil {
		log.Fatal("Can't get user by id", id)
		return Employee{}, err
	}
	return employee, err
}

func GetEmployees() ([]Employee, error) {
	query := "SELECT * FROM employees"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal("Can't get customers")
		return nil, err
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var employee Employee
		err := rows.Scan(
			&employee.ID,
			&employee.FirstName,
			&employee.LastName,
			&employee.Patronymic,
			&employee.User.UserName,
			&employee.Storage.Region)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, err
}

func UpdateEmployee(id uint64) error {
	_, err := db.Exec("UPDATE employee SET name=$1, email=$2 WHERE id=$3", name, email, id)
	if err != nil {
		log.Fatal("Can't update employee")
		return err
	}
	return nil
}

func DeleteEmployee(id uint64) error {
	deleteQuery := "DELETE FROM employyes WHERE id=?"
	_, err := db.Exec(deleteQuery, id)
	if err != nil {
		log.Fatal("Can't remove cookie")
	}
	return nil
}
