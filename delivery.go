package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Delivery struct {
	Date              Date           `xml:"Date,omitempty"`
	DeliveryMethod    DeliveryMethod `xml:"DeliveryMethod,omitempty"`
	TimeInDays        int            `xml:"TimeInDays,omitempty"`
	FromStock         bool           `xml:"FromStock,omitempty"`
	DropShip          bool           `xml:"DropShip,omitempty"`
	CountryOfAssembly string         `xml:"CountryOfAssembly,omitempty"`
}

func (d Delivery) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(d, e, start)
}

func (d Delivery) IsEmpty() bool {
	return zero.IsZero(d)
}
