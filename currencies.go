package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Currencies []Currency

type Currency struct {
	// Attributes
	Code   string `xml:"code,attr"`
	Active bool   `xml:"active,attr,omitempty"`

	Description                string    `xml:"Description,omitempty"`
	VariableExchangeRate       bool      `xml:"VariableExchangeRate,omitempty"`
	PrecisionRates             float64   `xml:"PrecisionRates,omitempty"`
	PrecisionAmounts           float64   `xml:"PrecisionAmounts,omitempty"`
	PrecisionPrices            float64   `xml:"PrecisionPrices,omitempty"`
	InEMU                      bool      `xml:"InEMU,omitempty"`
	InEMUSince                 Date      `xml:"InEMUSince,omitempty"`
	SWIFTCurrency              string    `xml:"SWIFTCurrency,omitempty"`
	CurrencyCharacter          string    `xml:"CurrencyCharacter,omitempty"`
	Factor                     float64   `xml:"Factor,omitempty"`
	GLPurchaseLoss             GLAccount `xml:"GLPurchaseLoss,omitempty"`
	GLPurchaseGain             GLAccount `xml:"GLPurchaseGain,omitempty"`
	GLSalesLoss                GLAccount `xml:"GLSalesLoss,omitempty"`
	GLSalesGain                GLAccount `xml:"GLSalesGain,omitempty"`
	GLUnfavourableRevaluations GLAccount `xml:"GLUnfavourableRevaluations,omitempty"`
	GLFavourableRevaluations   GLAccount `xml:"GLFavourableRevaluations,omitempty"`
}

func (c Currency) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c Currency) IsEmpty() bool {
	return zero.IsZero(c)
}
