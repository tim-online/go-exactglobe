package globe

type GLAccountCategory struct {
	// Attributes
	Number int    `xml:"number,attr"`
	Code   string `xml:"code,attr"`

	Description       string   `xml:"Description,omitempty"`
	MultiDescriptions []string `xml:"MultiDescriptions>MultiDescription,omitempty"`
}
