package globe

type TaxCode struct {
	// Attributes
	Code    string `xml:"code,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`    // { V | I | B }
	VatType string `xml:"vattype,attr,omitempty"` // { E | I | N }
	TaxType string `xml:"taxtype,attr,omitempty"` // { V | W }

	// Description          string          `xml:"Description,omitempty"`
	MultiDescriptions []string `xml:"Description,omitempty"`

	// <Included> (0,1)
	// <Percentage> (0,1)
	// <Charged> (0,1)
	// <UseCashSystem> (0,1)
	// <VATExemption> (0,1)
	// <ExtraDutyPercentage> (0,1)
	// <GLToPay> (0,1)
	// <GLToClaim> (0,1)
	// <GLToPaySuspense> (0,1)
	// <GLToClaimSuspense> (0,1)
	// <Value> (0,1)
	// <Creditor> (0,1)
	// <PaymentPeriod> (0,1)
	// <VATListing> (0,1)
	// <VATBoxLink> (0,+)
	// <NonDeductiblePercentage> (0,1)
	// <GLNonDeductibleAccount> (0,1)
}
