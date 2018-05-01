package globe

type BankAccounts []BankAccount

type BankAccount struct {
	// Attributes
	Code    string `xml:"code,attr,omitempty"`
	Default bool   `xml:"default,attr"`

	BankAccountType   BankAccountType `xml:"BankAccountType,omitempty"`
	Currency          Currency        `xml:"Currency,omitempty"`
	Bank              Bank            `xml:"Bank,omitempty"`
	SDDMandates       SDDMandates     `xml:"SDDMandates,omitempty"`
	Branch            string          `xml:"Branch,omitempty"`
	AccountNumberBank string          `xml:"AccountNumberBank,omitempty"`
	Address           Address         `xml:"Address,omitempty"`
	Contact           Contact         `xml:"Contact,omitempty"`
}
