package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

// An amount in a foreign currency.
type ForeignAmount struct {
	// Attributes
	Type string `xml:"type,attr,omitempty"` // { S | P }

	Currency Currency `xml:"Currency"`
	Debit    float64  `xml:"Debit,omitempty"`
	Credit   float64  `xml:"Credit,omitempty"`
	Value    float64  `xml:"Value,omitempty"`
	Rate     float64  `xml:"Rate"`
	VAT      VAT      `xml:"VAT,omitempty"`
}

func (f ForeignAmount) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(f, e, start)
}

func (f ForeignAmount) IsEmpty() bool {
	return zero.IsZero(f)
}
