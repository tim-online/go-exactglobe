package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type VATs []VAT

type VAT struct {
	// Attributes
	Code    string `xml:"code,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`    // { V | I | B }
	VatType string `xml:"vattype,attr,omitempty"` // { E | I | N }
	TaxType string `xml:"taxtype,attr,omitempty"` // { V | W }

	Description             string            `xml:"description,attr,omitempty"`
	MultiDescriptions       MultiDescriptions `xml:"MultiDescriptions,omitempty"`
	Included                bool              `xml:"Included,omitempty"`
	Percentage              float64           `xml:"Percentage,omitempty"`
	Charged                 bool              `xml:"Charged,omitempty"`
	UseCashSystem           bool              `xml:"UseCashSystem,omitempty"`
	VATExemption            bool              `xml:"VATExemption,omitempty"`
	ExtraDutyPercentage     interface{}       `xml:"ExtraDutyPercentage,omitempty"`
	GLToPay                 *GLAccount        `xml:"GLToPay,omitempty"`
	GLToClaim               *GLAccount        `xml:"GLToClaim,omitempty"`
	GLToPaySuspense         *GLAccount        `xml:"GLToPaySuspense,omitempty"`
	GLToClaimSuspense       *GLAccount        `xml:"GLToClaimSuspense,omitempty"`
	Value                   float64           `xml:"Value,omitempty"`
	Creditor                *Creditor         `xml:"Creditor,omitempty"`
	PaymentPeriod           string            `xml:"PaymentPeriod,omitempty"`
	VATListing              interface{}       `xml:"VATListing,omitempty"`
	VATBoxLink              VATBoxLink        `xml:"VATBoxLink,omitempty"`
	NonDeductiblePercentage float64           `xml:"NonDeductiblePercentage,omitempty"`
	GLNonDeductibleAccount  *GLAccount        `xml:"GLNonDeductibleAccount,omitempty"`
}

func (v VAT) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(v, e, start)
}

func (v VAT) IsEmpty() bool {
	return zero.IsZero(v)
}
