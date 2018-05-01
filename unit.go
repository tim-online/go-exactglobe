package globe

type Unit struct {
	// Attributes
	Unit string `xml:"unit,attr"`
	Type string `xml:"type,attr,omitempty"` // { L | W | T | O }

	// Description string `xml:"Description,omitempty"`
	MultiDescriptions []string `xml:"Description,omitempty"`
}
