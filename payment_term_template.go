package globe

type PaymentTermTemplates []PaymentTermTemplate

type PaymentTermTemplate struct {
	// Attributes
	Code          string `xml:"code"`
	Type          string `xml:"type,attr,omitempty"`        // { N | K | C | T | Q | Z | Y | R | P | S | D | F | M | L | E | A | B | G | H | J | W | O | X | U }
	PaymentType   string `xml:"paymentType,attr,omitempty"` // { B | K | R | C | W | I | O | F | P | S | V | Y | X | E | H | M | N }
	PaymentMethod string `xml:"paymentMethod,attr,omitempty"`

	DaysToPayment     int            `xml:"DaysToPayment"`
	Percentage        float64        `xml:"Percentage"`
	DueDate           Date           `xml:"DueDate,omitempty"`
	MaturityDays      int            `xml:"MaturityDays"`
	OwnBankAccount    OwnBankAccount `xml:"OwnBankAccount"`
	OwnBankAccountRef string         `xml:"OwnBankAccountRef"`
}
