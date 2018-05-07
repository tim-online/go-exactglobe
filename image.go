package globe

type Image struct {
	// Attributes
	ID string `xml:"ID,attr,omitempty"`

	Name       string `xml:"Name"`
	BinaryData []byte `xml:"BinaryData,omitempty"`
}
