package customers

import (
	. "../../models/customer"
	"log"
)

func AddOrganizationCustomer(organization *Organization) error {
	query := "INSERT INTO organizations(customer_id, organization_name, organization_address, bank_details, contact_person) VALUES(?,?,?,?,?,?)"

	_, err := db.Exec(query,
		organization.Customer.ID,
		organization.OrganizationName,
		organization.OrganizationAddress,
		organization.BankDetails.ID,
		organization.ContactPerson)

	if err != nil {
		log.Fatal("Can't add organization customer in DB")
		return err
	}
	return nil
}
