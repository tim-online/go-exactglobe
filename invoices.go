package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	errors "github.com/tim-online/go-errors"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Invoices []Invoice

type Invoice struct {
	// Attributes
	Type     InvoiceType `xml:"type,attr"`
	Code     string      `xml:"code,attr"`
	Number   int         `xml:"number,attr"`
	Original string      `xml:"original,attr"`
	Sequence int         `xml:"sequence,attr"`

	InvoiceLines InvoiceLines `xml:"InvoiceLine"`
}

func (i Invoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(i, e, start)
}

func (i Invoice) IsEmpty() bool {
	return zero.IsZero(i)
}

// V=Sales invoice, I=Commision invoice, B=Direct Invoice
type InvoiceType string

func (it InvoiceType) Validate() error {
	valid := []InvoiceType{"V", "I", "B"}

	if it == "" {
		return nil
	}

	for _, v := range valid {
		if it == v {
			return nil
		}
	}

	return errors.Errorf("Invalid InvoiceType '%s'", it)
}
