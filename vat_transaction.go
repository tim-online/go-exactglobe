package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type VATTransaction struct {
	// Attributes
	Code string `xml:"code,attr,omitempty"`

	VATBaseAmount         float64 `xml:"VATBaseAmount,omitempty"`
	VATBaseAmountFC       float64 `xml:"VATBaseAmountFC,omitempty"`
	OfficialAmountDC      float64 `xml:"OfficialAmountDC,omitempty"`
	OfficialBasisDC       float64 `xml:"OfficialBasisDC,omitempty"`
	OfficialExchangeRate  float64 `xml:"OfficialExchangeRate,omitempty"`
	VATPercentage         float64 `xml:"VATPercentage,omitempty"`
	VATForeignBaseAmount  float64 `xml:"VATForeignBaseAmount,omitempty"`
	VATNumber             string  `xml:"VATNumber,omitempty"`
	ReportNumberVATReturn string  `xml:"ReportNumberVATReturn,omitempty"`
	ReportNumberCTListing string  `xml:"ReportNumberCTListing,omitempty"`
}

func (v VATTransaction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(v, e, start)
}

func (v VATTransaction) IsEmpty() bool {
	return zero.IsZero(v)
}
