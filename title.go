package globe

type Title struct {
	// Attributes
	Code string `xml:"code,attr"`

	Description  string `xml:"Description"`
	Salutation   string `xml:"Salutation,omitempty"`
	Abbreviation string `xml:"Abbreviation,omitempty"`
}
