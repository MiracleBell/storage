package provides

import . ".."

type Provider struct {
	ID                  uint64
	organizationName    string
	organizationAddress string
	contactPerson       string
	products            []Goods
}
