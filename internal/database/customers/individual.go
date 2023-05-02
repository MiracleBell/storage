package customers

import (
	. "../../models/customer"
	"log"
)

func AddIndividualCustomer(individual *Individual) error {
	query := "INSERT INTO individuals(customer_id, first_name, last_name, patronymic, delovery_address, phone) VALUES(?,?,?,?,?,?)"

	_, err := db.Exec(query,
		individual.Customer.ID,
		individual.FirstName,
		individual.LastName,
		individual.Patronymic,
		individual.DeliveryAddress,
		individual.Phone)

	if err != nil {
		log.Fatal("Can't add individual customer in DB")
		return err
	}
	return nil
}
