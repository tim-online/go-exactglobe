package globe

import (
	"encoding/xml"

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

	Division      Division      `xml:"Division"`
	DocumentDate  Date          `xml:"DocumentDate"` // gbkmut.docdate
	FinYear       FinYear       `xml:"FinYear"`
	FinPeriod     FinPeriod     `xml:"FinPeriod"`
	Date          Date          `xml:"Date"`
	Journal       Journal       `xml:"Journal"`
	Amount        float64       `xml:"Amount"`
	FreeFields    FreeFields    `xml:"FreeFields"`
	Documents     Documents     `xml:"Documents"`
	FinEntryLine  FinEntryLine  `xml:"FinEntryLine"`
	PaymentTerms  PaymentTerms  `xml:"PaymentTerms"`
	BankStatement BankStatement `xml:"BankStatement"`
}

func (f FinEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(f, e, start)
}
