package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Unit struct {
	// Attributes
	Unit string `xml:"unit,attr"`
	Type string `xml:"type,attr,omitempty"` // { L | W | T | O }

	// Description string `xml:"Description,omitempty"`
	MultiDescriptions []string `xml:"Description,omitempty"`
}

func (u Unit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(u, e, start)
}

func (u Unit) IsEmpty() bool {
	return zero.IsZero(u)
}
