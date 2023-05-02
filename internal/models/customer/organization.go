package customer

import . "../detail"

type Organization struct {
	Customer            Customer
	OrganizationName    string
	OrganizationAddress string
	ContactPerson       string
	BankDetails         BankDetails
}
