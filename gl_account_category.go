package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type GLAccountCategory struct {
	// Attributes
	Number int    `xml:"number,attr"`
	Code   string `xml:"code,attr"`

	Description       string   `xml:"Description,omitempty"`
	MultiDescriptions []string `xml:"MultiDescriptions>MultiDescription,omitempty"`
}

func (g GLAccountCategory) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(g, e, start)
}

func (g GLAccountCategory) IsEmpty() bool {
	return zero.IsZero(g)
}
