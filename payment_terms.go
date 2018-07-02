package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	errors "github.com/tim-online/go-errors"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type PaymentTerms []PaymentTerm

// docs/XML-Schema.html#ELEMENT_PaymentTerm

type PaymentTerm struct {
	// Attributes
	Type          PaymentTermType          `xml:"type,attr"`
	Status        PaymentTermStatus        `xml:"status,attr"`
	Entry         int                      `xml:"entry,attr,omitempty"`
	ID            string                   `xml:"id,attr,omitempty"`
	MatchID       string                   `xml:"matchID,attr,omitempty"`
	PaymentMethod PaymentTermPaymentMethod `xml:"paymentMethod,attr,omitempty"`
	StatementType PaymentTermStatementType `xml:"statementType,attr,omitempty"`
	PaymentType   PaymentTermPaymentType   `xml:"paymentType,attr,omitempty"`
	TermType      PaymentTermTermType      `xml:"termType,attr,omitempty"`

	// type = token: { A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z }
	// status = token: { A | C | D | J | P | R | U | V }
	// [ entry = integer ]
	// [ ID = string ]
	// [ matchID = string ]

	// [ paymentMethod = token: { C | D | T | B | V | E | M | A | I | S | N } ]
	// [ statementType = token: { B | C | F | H | J | K | M | Q | U | W } ]
	// [ paymentType = token: { A | B | C | E | F | H | I | K | O | P | R | S | V | W | X | Y } ]
	// [ termType = token: { C | D | N | P | S | V | W } ]>

	Description          string           `xml:"Description"`
	GLAccount            GLAccount        `xml:"GLAccount,omitempty"`
	GLOffset             GLAccount        `xml:"GLOffset"`
	OwnBankAccount       OwnBankAccount   `xml:"OwnBankAccount,omitempty"`
	BankAccount          BankAccount      `xml:"BankAccount,omitempty"`
	MandateReference     interface{}      `xml:"MandateReference,omitempty"`
	OffsetBankName       string           `xml:"OffsetBankName,omitempty"`
	OffsetBankCountry    string           `xml:"OffsetBankCountry,omitempty"`
	Debtor               Debtor           `xml:"Debtor,omitempty"`
	Creditor             Creditor         `xml:"Creditor,omitempty"`
	TransactionNumber    string           `xml:"TransactionNumber,omitempty"`
	Amount               Amount           `xml:"Amount"`
	ForeignAmount        ForeignAmount    `xml:"ForeignAmount"`
	PaymentCondition     PaymentCondition `xml:"PaymentCondition"`
	DaysToPayment        int              `xml:"DaysToPayment,omitempty"`
	Percentage           float64          `xml:"Percentage"`
	Reference            string           `xml:"Reference"`
	YourRef              string           `xml:"YourRef"`
	OrderNumber          string           `xml:"OrderNumber,omitempty"`
	InvoiceNumber        string           `xml:"InvoiceNumber,omitempty"`
	InvoiceDate          Date             `xml:"InvoiceDate"`
	InvoiceDueDate       Date             `xml:"InvoiceDueDate"`
	ProcessingDate       Date             `xml:"ProcessingDate,omitempty"`
	ReportingDate        Date             `xml:"ReportingDate,omitempty"`
	OfficialAmountDC     float64          `xml:"OfficialAmountDC,omitempty"`
	OfficialBasisDC      float64          `xml:"OfficialBasisDC,omitempty"`
	OfficialExchangeRate float64          `xml:"OfficialExchangeRate,omitempty"`
	TaxInvoiceDate       Date             `xml:"TaxInvoiceDate,omitempty"`
	Resource             Resource         `xml:"Resource,omitempty"`
	Journalization       Journalization   `xml:"Journalization"`
	Warehouse            Warehouse        `xml:"Warehouse,omitempty"`
	Approving            Approving        `xml:"Approving,omitempty"`
	IsBlocked            bool             `xml:"IsBlocked,omitempty"`
}

func (p PaymentTerm) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(p, e, start)
}

func (p PaymentTerm) IsEmpty() bool {
	return zero.IsZero(p)
}

// Iota not in docs
type PaymentTermType string

func (p PaymentTermType) Validate() error {
	valid := []PaymentTermType{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if p == "" {
		return nil
	}

	for _, v := range valid {
		if p == v {
			return nil
		}
	}

	return errors.Errorf("Invalid PaymentTermType '%s'", p)
}

// Iota not in docs
type PaymentTermStatus string

func (p PaymentTermStatus) Validate() error {
	valid := []PaymentTermStatus{"A", "C", "D", "J", "P", "R", "U", "V"}

	if p == "" {
		return nil
	}

	for _, v := range valid {
		if p == v {
			return nil
		}
	}

	return errors.Errorf("Invalid PaymentTermStatus '%s'", p)
}

// Iota not in docs
type PaymentTermPaymentMethod string

func (p PaymentTermPaymentMethod) Validate() error {
	valid := []PaymentTermPaymentMethod{"C", "D", "T", "B", "V", "E", "M", "A", "I", "S", "N"}

	if p == "" {
		return nil
	}

	for _, v := range valid {
		if p == v {
			return nil
		}
	}

	return errors.Errorf("Invalid PaymentTermPaymentMethod '%s'", p)
}

// Iota not in docs
type PaymentTermStatementType string

func (p PaymentTermStatementType) Validate() error {
	valid := []PaymentTermStatementType{"B", "C", "F", "H", "J", "K", "M", "Q", "U", "W"}

	if p == "" {
		return nil
	}

	for _, v := range valid {
		if p == v {
			return nil
		}
	}

	return errors.Errorf("Invalid PaymentTermStatementType '%s'", p)
}

// Iota not in docs
type PaymentTermPaymentType string

func (p PaymentTermPaymentType) Validate() error {
	valid := []PaymentTermPaymentType{"A", "B", "C", "E", "F", "H", "I", "K", "O", "P", "R", "S", "V", "W", "X", "Y"}

	if p == "" {
		return nil
	}

	for _, v := range valid {
		if p == v {
			return nil
		}
	}

	return errors.Errorf("Invalid PaymentTermPaymentType '%s'", p)
}

// Iota not in docs
type PaymentTermTermType string

func (p PaymentTermTermType) Validate() error {
	valid := []PaymentTermTermType{"C", "D", "N", "P", "S", "V", "W"}

	if p == "" {
		return nil
	}

	for _, v := range valid {
		if p == v {
			return nil
		}
	}

	return errors.Errorf("Invalid PaymentTermTermType '%s'", p)
}
