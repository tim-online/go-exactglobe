package globe

type ExtraText struct {
	// Attributes
	Code string `xml:"code,attr"`

	Description       string            `xml:"Description,omitempty"`
	MultiDescriptions MultiDescriptions `xml:"MultiDescriptions,omitempty"`
}
