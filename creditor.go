package globe

type Creditor struct {
	// Attributes
	Code   string `xml:"code,attr"`
	Number string `xml:"number,attr,omitempty"`
	Type   string `xml:"type,attr,omitempty"` // { A | B | C | D | E | L | N | P | R | S }
	ID     string `xml:"ID,attr,omitempty"`

	Name             string           `xml:"Name,omitempty"`
	Currency         Currency         `xml:"Currency,omitempty"`
	SecurityLevel    int              `xml:"SecurityLevel,omitempty"`
	BankAccounts     BankAccounts     `xml:"BankAccounts,omitempty"`
	AutoMatching     bool             `xml:"AutoMatching,omitempty"`
	GLOffset         GLAccount        `xml:"GLOffset,omitempty"`
	GLCentralization GLAccount        `xml:"GLCentralization,omitempty"`
	ExternalCode     string           `xml:"ExternalCode,omitempty"`
	CreditLine       float64          `xml:"CreditLine,omitempty"`
	Discount         float64          `xml:"Discount,omitempty"`
	CustomerCode     string           `xml:"CustomerCode,omitempty"`
	AccountCategory  AccountCategory  `xml:"AccountCategory,omitempty"`
	PriceList        PriceList        `xml:"PriceList,omitempty"`
	PaymentCondition PaymentCondition `xml:"PaymentCondition,omitempty"`
	VATNumber        string           `xml:"VATNumber,omitempty"`
	PaymentMethod    PaymentMethod    `xml:"PaymentMethod,omitempty"`
}
