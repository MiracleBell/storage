package provides

import . "../goods"
import . "../detail"

type Provider struct {
	ID                  uint64
	BankDetails         BankDetails
	OrganizationName    string
	OrganizationAddress string
	ContactPerson       string
	Products            []Goods
}
