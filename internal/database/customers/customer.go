package customers

import (
	. "../"
	. "../../models/customer"
	"log"
)

var db, _ = OpenConnection()

func AddNewCustomer(customer *Customer) error {
	query := "INSERT INTO customers(customer_type, email, created_at, updated_at) VALUES(?,?,?,?)"

	result, err := db.Exec(query,
		customer.CustomerType,
		customer.Email,
		customer.CreatedAt,
		customer.UpdatedAt)

	if err != nil {
		log.Fatal("Can't add user in DB")
		return err
	}
	id, _ := result.LastInsertId()
	customer.ID = uint64(id)
	return nil
}

func GetCustomerById(id uint64) (Customer, error) {
	query := "SELECT * FROM users WHERE id=?"

	var customer Customer
	err := db.QueryRow(query, id).Scan(
		&customer.ID,
		&customer.CustomerType,
		&customer.Email,
		&customer.CreatedAt,
		&customer.UpdatedAt)

	if err != nil {
		log.Fatal("Can't get customers by id", id)
		return Customer{}, err
	}
	return customer, err
}

func GetCustomers() ([]Customer, error) {
	query := "SELECT * FROM users"

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal("Can't get customers")
		return nil, err
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.ID, &customer.CustomerType, &customer.Email)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, err
}
