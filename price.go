package globe

type Price struct {
	// Attributes
	Type string `xml:"type,attr,omitempty"`

	Float float64 `xml:",chardata"`
}
