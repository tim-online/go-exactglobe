package globe

// bnkacc.accncd
type BankAccountType struct {
	// Attributes
	Code      string `xml:"code,attr,omitempty"`
	CheckType string `xml:"checktype,omitempty"` //  { E | G | B | U | T | R | X | D | F | S | I | Y | Z | K | O | V | W | C | J | A | H | P | L | M | 0 | 1 | 2 | 4 | 5 | N }

	Description          string               `xml:"Description"`
	BankAccountMask      string               `xml:"BankAccountMask,omitempty"`
	BankAccountCheckType BankAccountCheckType `xml:"BankAccountCheckType,omitempty"`
}
