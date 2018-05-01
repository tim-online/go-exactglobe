package globe

type AccountCategory struct {
	// Attributes
	Code string `xml:"code,attr"`

	Description string `xml:"Description,omitempty"`
	Type        string `xml:"Type,omitempty"`
}
