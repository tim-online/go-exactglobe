package globe

// docs/XML-Schema.html#ELEMENT_Costcenter_Categories

type CostcenterCategories []CostcenterCategory

type CostcenterCategory struct {
	// Attributes
	Number int    `xml:"number,attr"`
	Code   string `xml:"code,attr"`

	Description       string `xml:"Description,omitempty"`
	CategoryGroupName string `xml:"CategoryGroupName"`
}
