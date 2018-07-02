package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type PaymentCondition struct {
	// Attributes
	Code         string      `xml:"code,attr"`
	Type         string      `xml:"type,attr,omitempty"` // { 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 }
	Type1        interface{} `xml:"type1,attr,omitempty"`
	Surcharge    string      `xml:"surcharge,attr,omitempty"` // { K | B }
	Method       string      `xml:"method,attr,omitempty"`    // { B | N }
	Invoicing    string      `xml:"invoicing,attr,omitempty"` // { B | N }
	Installments bool        `xml:"installments,attr,omitempty"`

	// Description          string          `xml:"Description,omitempty"`
	MultiDescriptions    []string             `xml:"Description,omitempty"`
	DaysToPayment        int                  `xml:"DaysToPayment,omitempty"`
	NumberOfMonths       int                  `xml:"NumberOfMonths,omitempty"`
	DayOfTheMonth        int                  `xml:"DayOfTheMonth,omitempty"`
	Term1                Term                 `xml:"Term_1,omitempty"`
	Term2                Term                 `xml:"Term_2,omitempty"`
	Term3                Term                 `xml:"Term_3,omitempty"`
	Assortment           Assortment           `xml:"Assortment,omitempty"`
	PaymentTermTemplates PaymentTermTemplates `xml:"PaymentTermTemplates,omitempty"`
}

func (pc PaymentCondition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(pc, e, start)
}

func (pc PaymentCondition) IsEmpty() bool {
	return zero.IsZero(pc)
}
