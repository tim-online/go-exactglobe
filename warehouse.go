package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Warehouse struct {
	// Attributes
	Code    string `xml:"code,attr"`
	Blocked bool   `xml:"bolecked,attr,omitempty"`

	Description string  `xml:"Description"`
	Address     Address `xml:"Address,omitempty"`
	DropShip    bool    `xml:"DropShip,omitempty"`
}

func (w Warehouse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(w, e, start)
}

func (w Warehouse) IsEmpty() bool {
	return zero.IsZero(w)
}
