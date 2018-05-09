package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type PriceList struct {
	// Attributes
	Code string `xml:"code,attr"`
	Type string `xml:"type,attr,omitempty"` // { S | P }

	// Description          string          `xml:"Description,omitempty"`
	MultiDescriptions       []string     `xml:"Description,omitempty"`
	Currency                Currency     `xml:"Currency,omitempty"`
	Country                 Country      `xml:"Country,omitempty"`
	Availability            Availability `xml:"Availability,omitempty"`
	AccountType             string       `xml:"AccountType,omitempty"`
	AccountStatus           string       `xml:"AccountStatus,omitempty"`
	AccountClassificationId string       `xml:"AccountClassificationId,omitempty"`
}

func (p PriceList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(p, e, start)
}

func (p PriceList) IsEmpty() bool {
	return zero.IsZero(p)
}
