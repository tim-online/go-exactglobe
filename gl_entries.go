package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type GLEntries []GLEntry

type GLEntry struct {
	// Attributes
	Entry  string        `xml:"entry,attr,omitempty"`
	Status GLEntryStatus `xml:"status,attr,omitempty"` // GLEntryStatus

	Division      Division      `xml:"Division,omitempty"`
	Description   string        `xml:"Description,omitempty"`
	Date          Date          `xml:"Date,omitempty"`
	DocumentDate  Date          `xml:"DocumentDate,omitempty"`
	Project       Project       `xml:"Project,omitempty"`
	Journal       Journal       `xml:"Journal,omitempty"`
	Costcenter    Costcenter    `xml:"Costcenter,omitempty"`
	Costunit      Costunit      `xml:"Costunit,omitempty"`
	Amount        Amount        `xml:"Amount,omitempty"`
	ForeignAmount ForeignAmount `xml:"ForeignAmount,omitempty"`
	FreeFields    FreeFields    `xml:"FreeFields,omitempty"`
	FinEntryLines FinEntryLines `xml:">FinEntryLine"`
	PaymentTerms  PaymentTerms  `xml:"PaymentTerms,omitempty"`
	BankStatement BankStatement `xml:"BankStatement,omitempty"`
}

func (g GLEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(g, e, start)
}

func (g GLEntry) IsEmpty() bool {
	return zero.IsZero(g)
}

type GLEntryStatus string
