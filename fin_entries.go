package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type FinEntries []FinEntry

// Header of a financial entry. Each financial entry has to be in balance and
// contains at least 2 detail lines (financial entry lines). The header contains
// data which are applicable for all financial entry lines which belongs to an
// header.
type FinEntry struct {
	// Attributes
	Entry    string `xml:"entry,attr"`              // Entry number of the financial entry.
	ID       string `xml:"ID,attr,omitempty"`       // ID to identify the Entry.
	Sequence int    `xml:"sequence,attr,omitempty"` // gbkmut.volgnr5

	Division      Division      `xml:"Division,omitempty"`
	DocumentDate  Date          `xml:"DocumentDate,omitempty"` // gbkmut.docdate
	FinYear       FinYear       `xml:"FinYear,omitempty"`
	FinPeriod     FinPeriod     `xml:"FinPeriod,omitempty"`
	Date          Date          `xml:"Date,omitempty"`
	Journal       Journal       `xml:"Journal,omitempty"`
	Amount        float64       `xml:"Amount,omitempty"`
	FreeFields    FreeFields    `xml:"FreeFields,omitempty"`
	Documents     Documents     `xml:"Documents,omitempty"`
	FinEntryLine  FinEntryLine  `xml:"FinEntryLine"`
	PaymentTerms  PaymentTerms  `xml:"PaymentTerms,omitempty"`
	BankStatement BankStatement `xml:"BankStatement,omitempty"`
}

func (f FinEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(f, e, start)
}

func (f FinEntry) IsEmpty() bool {
	return zero.IsZero(f)
}
