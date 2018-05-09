package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type WarehouseLocation struct {
	// Attributes
	Code string `xml:"code,attr"`

	Warehouse   Warehouse `xml:"Warehouse"`
	Description string    `xml:"Description"`
}

func (w WarehouseLocation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(w, e, start)
}

func (w WarehouseLocation) IsEmpty() bool {
	return zero.IsZero(w)
}
