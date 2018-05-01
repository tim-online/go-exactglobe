package globe

type ItemAccounts []ItemAccount

type ItemAccount struct {
	// Attributes
	Default bool `xml:"default"`

	Account            Account      `xml:"Account"`
	ItemCode           string       `xml:"ItemCode"`
	ItemCodeAccount    string       `xml:"ItemCodeAccount,omitempty"`
	SupplierPreference int          `xml:"SupplierPreference,omitempty"`
	EANCode            string       `xml:"EANCode,omitempty"`
	Purchase           Purchase     `xml:"Purchase,omitempty"`
	Manufacturer       Manufacturer `xml:"Manufacturer,omitempty"`
	Delivery           Delivery     `xml:"Delivery,omitempty"`
	DropShip           bool         `xml:"DropShip,omitempty"`
	CountryOfAssembly  string       `xml:"CountryOfAssembly,omitempty"`
}
