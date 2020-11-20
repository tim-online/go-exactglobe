package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	errors "github.com/tim-online/go-errors"
	"github.com/tim-online/go-exactglobe/omitempty"
)

// An amount in the default currency.
type Amount struct {
	// Attributes
	Type AmountType `xml:"type,attr,omitempty"`

	Currency    Currency `xml:"Currency"`
	Debit       float64  `xml:"Debit"`
	VAT         VAT      `xml:"VAT,omitempty"`
	Credit      float64  `xml:"Credit"`
	Value       float64  `xml:"Value,omitempty"`
	TaxCode2    TaxCode  `xml:"TaxCode2,omitempty"`
	TaxCode3    TaxCode  `xml:"TaxCode3,omitempty"`
	TaxCode4    TaxCode  `xml:"TaxCode4,omitempty"`
	TaxCode5    TaxCode  `xml:"TaxCode5,omitempty"`
	LineCharges float64  `xml:"LineCharges,omitempty"`
}

func (a Amount) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(a, e, start)
}

func (a Amount) IsEmpty() bool {
	return zero.IsZero(a)
}

type AmountType string

func (a AmountType) Validate() error {
	valid := []AmountType{"S", "P"}

	if a == "" {
		return nil
	}

	for _, v := range valid {
		if a == v {
			return nil
		}
	}

	return errors.Errorf("Invalid AmountType '%s'", a)
}
