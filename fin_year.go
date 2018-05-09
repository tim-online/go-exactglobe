package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type FinYear struct {
	// Attributes
	Number int `xml:"number,attr,omitempty"`

	FinPeriod FinPeriod `xml:"FinPeriod,omitempty"`
}

func (f FinYear) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(f, e, start)
}

func (f FinYear) IsEmpty() bool {
	return zero.IsZero(f)
}
