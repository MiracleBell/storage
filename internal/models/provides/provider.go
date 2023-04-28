package provides

import . "../goods"
import . "../detail"

type Provider struct {
	ID                  uint64
	bankDetails         BankDetails
	organizationName    string
	organizationAddress string
	contactPerson       string
	products            []Goods
}
