package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type GLAccounts []GLAccount

type GLAccount struct {
	// Attributes
	Code                  string `xml:"code,attr"`
	Type                  string `xml:"type,attr,omitempty"`         // { B | W }
	Subtype               string `xml:"subtype,attr,omitempty"`      // { A | B | C | D | G | H | J | K | N | S | T | V }
	Side                  string `xml:"side,attr,omitempty"`         // { D | C | G } ]
	Presentation          string `xml:"presentation,attr,omitempty"` // { J | L | N } ]
	Blocked               bool   `xml:"blocked,attr,omitempty"`
	Inflationadjustment   string `xml:"inflationadjustment,attr,omitempty,omitempty"` // { N | I | A }
	Invoiceregtype        string `xml:"invoiceregtype,attr,omitempty"`                // { I | V | N }
	Purchasevatreturntype string `xml:"purchasevatreturntype,attr,omitempty"`         // { A | D | G | I | O }
	Rewardtype            string `xml:"rewardtype,attr,omitempty"`                    // { C | E | K | V | N }

	// Description                string            `xml:"Description,omitempty"`
	MultiDescriptions          []string          `xml:"Description,omitempty"`
	IsCostcenterAccount        bool              `xml:"IsCostcenterAccount,omitempty"`
	IsQuantityAccount          bool              `xml:"IsQuantityAccount,omitempty"`
	IsItemAccount              bool              `xml:"IsItemAccount,omitempty"`
	IsResourceAccount          bool              `xml:"IsResourceAccount,omitempty"`
	IsProjectAccount           bool              `xml:"IsProjectAccount,omitempty"`
	IsAssetAccount             bool              `xml:"IsAssetAccount,omitempty"`
	IsPersonalAccount          bool              `xml:"IsPersonalAccount,omitempty"`
	IsLinked                   bool              `xml:"IsLinked,omitempty"`
	Revalue                    bool              `xml:"Revalue,omitempty"`
	Compress                   bool              `xml:"Compress,omitempty"`
	Costcenter                 Costcenter        `xml:"Costcenter,omitempty"`
	Costunit                   Costunit          `xml:"Costunit,omitempty"`
	VAT                        VAT               `xml:"VAT,omitempty"`
	GLAccountCategory          GLAccountCategory `xml:"GLAccountCategory,omitempty"`
	GLAssociate                *GLAccount        `xml:"GLAssociate,omitempty"`
	GLAdjustmentDebit          *GLAccount        `xml:"GLAdjustmentDebit,omitempty"`
	GLAdjustmentCredit         *GLAccount        `xml:"GLAdjustmentCredit,omitempty"`
	GLVATNonDeductable         *GLAccount        `xml:"GLVATNonDeductable,omitempty"`
	GLCorporate                *GLAccount        `xml:"GLCorporate,omitempty"`
	AccountReportCategory      int               `xml:"AccountReportCategory,omitempty"`
	PercentageVATNonDeductable float64           `xml:"PercentageVATNonDeductable,omitempty"`
	CostcenterLinks            Costcenters       `xml:"CostcenterLinks,omitempty"`
	CostunitLinks              Costunits         `xml:"CostunitLinks,omitempty"`
	Notes                      string            `xml:"Notes,omitempty"`
}

func (g GLAccount) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(g, e, start)
}

func (g GLAccount) IsEmpty() bool {
	return zero.IsZero(g)
}
