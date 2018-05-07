package globe

type Language struct {
	// Attributes
	Code string `xml:"code,attr"`

	Description string `xml:"Description,omitempty"`
}
