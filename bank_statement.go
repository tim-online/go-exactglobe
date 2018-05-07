package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type BankStatement struct {
	// Attributes
	Number string `xml:"number,attr,omitempty"`

	Date              Date               `xml":Date,omitempty`
	GLOffset          GLAccount          `xml:"GLOffset,omitempty"`
	BankStatementLine BankStatementLines `xml:">BankStatementLine"`
}

func (b BankStatement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(b, e, start)
}

func (b BankStatement) IsEmpty() bool {
	return zero.IsZero(b)
}

type BankStatementLines []BankStatementLine

type BankStatementLine struct {
	// Attributes
	Type          string `xml:"type,attr"`             // { A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z }
	Status        string `xml:"status,attr,omitempty"` // { A | C | D | J | P | R | U | V }
	Entry         string `xml:"entry,attr,omitempty"`
	ID            string `xml:"ID,attr,omitempty"`
	LineNo        int    `xml:"lineNo,attr,omitempty"`
	PaymentMethod string `xml:"paymentMethod,attr,omitempty"` // { C | D | T }
	StatementType string `xml:"statementType,attr,omitempty"` // { B | C | F | H | J | K | M | Q | U | W }
	PaymentType   string `xml:"paymentType,attr,omitempty"`   // { B | C | E | F | H | I | K | O | P | R | S | V | W | X | Y }
	TermType      string `xml:"termType,attr,omitempty"`      // { C | D | N | P | S | V | W }

	Description       string         `xml:"Description"`
	ValueDate         Date           `xml:"ValueDate"`
	ReportingDate     Date           `xml:"ReportingDate"`
	StatementDate     Date           `xml:"StatementDate"`
	GLAccount         GLAccount      `xml:"GLAccount,omitempty"`
	OwnBankAccount    OwnBankAccount `xml:"OwnBankAccount,omitempty"`
	BankAccount       BankAccount    `xml:"BankAccount,omitempty"`
	Debtor            Debtor         `xml:"Debtor,omitempty"`
	Creditor          Creditor       `xml:"Creditor,omitempty"`
	TransactionNumber string         `xml:"TransactionNumber,omitempty"`
	Amount            Amount         `xml:"Amount"`
	ForeignAmount     ForeignAmount  `xml:"ForeignAmount"`
	Reference         string         `xml:"Reference"`
	YourRef           string         `xml:"YourRef"`
	InvoiceNumber     int            `xml:"InvoiceNumber"`
	Resource          Resource       `xml:"Resource,omitempty"`
	Journalization    Journalization `xml:"Journalization,omitempty"`
	IsBlocked         bool           `xml:"IsBlocked,omitempty"`
	Warehouse         Warehouse      `xml:"Warehouse,omitempty"`
	Approving         Approving      `xml:"Approving,omitempty"`
	PaymentTermIDs    []string       `xml:"PaymentTermIDs,omitempty"`
	GLOffset          GLAccount      `xml:"GLOffset,omitempty"`
}
