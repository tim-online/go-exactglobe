package globe

type Manufacturer struct {
	// Attributes
	Code   string `xml:"code,attr"`
	Type   string `xml:"type,attr,omitempty"`
	Status string `xml:"status,attr,omitempty"`

	Name string `xml:"Name"`
}
