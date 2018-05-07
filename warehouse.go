package globe

type Warehouse struct {
	// Attributes
	Code    string `xml:"code,attr"`
	Blocked bool   `xml:"bolecked,attr,omitempty"`

	Description string  `xml:"Description"`
	Address     Address `xml:"Address,omitempty"`
	DropShip    bool    `xml:"DropShip,omitempty"`
}
