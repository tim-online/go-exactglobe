package globe

type OwnBankAccount struct {
	// Attributes
	Code string `xml:"code,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"` // { C | K | P | R | U }

	Description        string    `xml:"Description"`
	Currency           Currency  `xml:"Currency"`
	Journal            Journal   `xml:"Journal,omitempty"`
	GLAccount          GLAccount `xml:"GLAccount"`
	GLPaymentInTransit GLAccount `xml:"GLPaymentInTransit"`
	Country            Country   `xml:"Country"`
	BankName           string    `xml:"BankName"`
	BankCreditor       string    `xml:"BankCreditor"`
	IBAN               string    `xml:"IBAN,omitempty"`
	BIC                string    `xml:"BIC,omitempty"`
}
