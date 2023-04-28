package customer

import . "../detail"

type Organization struct {
	customer            Customer
	organizationName    string
	organizationAddress string
	contactPerson       string
	bankDetails         BankDetails
}
