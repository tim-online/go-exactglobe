package globe

type IncoTermCode struct {
	// Attributes
	Code string `xml:"code,attr"`

	Description      string           `xml:"Description"`
	IncoTermProperty IncoTermProperty `xml:"IncoTermProperty,omitempty"`
}
