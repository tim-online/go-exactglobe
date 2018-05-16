package globe

type Categories []Category

type Category struct {
	// Attributs
	Number int    `xml:"number,attr"`
	Code   string `xml:"code,attr"`

	Description       string `xml:"Description,omitempty"`
	CategoryGroupName string `xml:"CategoryGroupName,omitempty"`
}
