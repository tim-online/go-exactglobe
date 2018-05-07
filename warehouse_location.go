package globe

type WarehouseLocation struct {
	// Attributes
	Code string `xml:"code,attr"`

	Warehouse   Warehouse `xml:"Warehouse"`
	Description string    `xml:"Description"`
}
