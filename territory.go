package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Territories []Territory

type Territory struct {
	// Attributes
	Code string `xml:"code,attr,omitempty"`

	// Description         string   `xml:"Description,omitempty"`
	MultiDescriptions []string `xml:"Description,omitempty"`
}

func (t Territory) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(t, enc, start)
}

func (t Territory) IsEmpty() bool {
	return zero.IsZero(t)
}
