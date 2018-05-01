package globe

type DeliveryMethod struct {
	// Attributes
	Code string `xml:"code,attr"`
	Type string `xml:"type,attr,omitempty"` // { N | C | P | O | E }

	// Description         string   `xml:"Description,omitempty"`
	MultiDescriptions   []string `xml:"Description,omitempty"`
	DirectShipping      bool     `xml:"DirectShipping,omitempty"`
	UseShippingManifest bool     `xml:"UseShippingManifest,omitempty"`
}
