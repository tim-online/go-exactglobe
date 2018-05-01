package globe

// bnkacc.bankcode
type Bank struct {
	// Attributes
	Code string `xml:"code,omitempty"`

	Name      string  `xml:"Name"`
	Country   Country `xml:"Country,omitempty"`
	IBAN      string  `xml:"IBAN,omitempty"`
	BIC       string  `xml:"BIC,omitempty"`
	SwiftCode string  `xml:"SwiftCode,omitempty"`
}
