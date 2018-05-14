package globe

type VATBoxLink struct {
	// Attributes
	Sign string `xml:"sign,attr,omitempty"` // { P | N }

	VATBox VATBox `xml:"VATBox"`
}
