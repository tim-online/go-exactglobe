package globe

type Nationality struct {
	// Attributes
	Code string `xml:"code,attr"`

	Description string `xml:"Description,omitempty"`
}
