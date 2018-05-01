package globe

type PaymentTerms []PaymentTerm

// docs/XML-Schema.html#ELEMENT_PaymentTerm

type PaymentTerm struct {
	// Attributes
	Number string `xml:"number,attr,omitempty"`

	Date     Date      `xml:"Date,omitempty"`
	GLOffset GLAccount `xml:"GLOffset,omitempty"`
}
