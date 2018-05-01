package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Division struct {
	// Attributes
	Code string `xml:"code,attr"`

	Currency    Currency `xml:"Currency,omitempty"`
	Description string   `xml:"Description,omitempty"`
}

func (d Division) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(d, e, start)
}

func (d Division) IsEmpty() bool {
	return zero.IsZero(d)
}
