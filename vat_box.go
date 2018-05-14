package globe

type VATBox struct {
	// Attributes
	Code     string `xml:"code,attr"`
	Type     string `xml:"type,attr"` // { A | I | O | R | T | V }
	Retrieve bool   `xml:"retrieve,attr,omitempty"`

	Description string  `xml:"Description,omitempty"`
	Country     Country `xml:"Country"`
}
