package globe

type Territories []Territory

type Territory struct {
	// Attributes
	Code string `xml:"code,attr,omitempty"`

	// Description         string   `xml:"Description,omitempty"`
	MultiDescriptions []string `xml:"Description,omitempty"`
}
