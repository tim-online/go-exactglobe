package globe

type Country struct {
	// Attributes
	Code string `xml:"code,attr"`

	Description string `xml:"Description,omitempty"`
}
