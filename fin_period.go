package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

// Start Date: YYYY-MM-DD. Time is optional here; if it must be included then
// enter the time in the format HH:MM:SS and seperate the two parts with a
// capital 'T': YYYY-MM-DDTHH:MM:SS. However, please note that when the time
// part is included validation of this tag will fail when using the XML
// validator (import will not fail, of course...).
type FinPeriod struct {
	// Attributes
	Number int `xml:"number,attr,omitempty"`

	DateStart Date `xml:"DateStart,omitempty"`
	DateEnd   Date `xml:"DateEnd,omitempty"`
}

func (f FinPeriod) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(f, e, start)
}

func (f FinPeriod) IsEmpty() bool {
	return zero.IsZero(f)
}
