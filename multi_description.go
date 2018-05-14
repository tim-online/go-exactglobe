package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type MultiDescriptions []MultiDescription

type MultiDescription struct {
	// Attributes
	Number int `xml:"Number,attr"`

	String string `xml:",chardata"`
}

func (m MultiDescription) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(m, e, start)
}

func (m MultiDescription) IsEmpty() bool {
	return zero.IsZero(m)
}
