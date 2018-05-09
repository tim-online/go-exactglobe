package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Selection struct {
	// Attributes
	Code string `xml:"code,attr"`

	Description string `xml:"Description,omitempty"`
}

func (s Selection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(s, e, start)
}

func (s Selection) IsEmpty() bool {
	return zero.IsZero(s)
}
