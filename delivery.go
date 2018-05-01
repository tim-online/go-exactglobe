package globe

type Delivery struct {
	Date              Date           `xml:"Date,omitempty"`
	DeliveryMethod    DeliveryMethod `xml:"DeliveryMethod,omitempty"`
	TimeInDays        int            `xml:"TimeInDays,omitempty"`
	FromStock         bool           `xml:"FromStock,omitempty"`
	DropShip          bool           `xml:"DropShip,omitempty"`
	CountryOfAssembly string         `xml:"CountryOfAssembly,omitempty"`
}
