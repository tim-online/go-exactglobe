package globe

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
