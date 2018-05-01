package globe

type Purchase struct {
	Price      Price   `xml:"Price"`
	Unit       Unit    `xml:"Unit,omitempty"`
	Package    string  `xml:"Package,omitempty"`
	SalesUnits float64 `xml:"SalesUnits,omitempty"`
	OrderSize  float64 `xml:"OrderSize,omitempty"`
}
