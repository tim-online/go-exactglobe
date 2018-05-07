package globe

type Selection struct {
	// Attributes
	Code string `xml:"code,attr"`

	Description string `xml:"Description,omitempty"`
}
