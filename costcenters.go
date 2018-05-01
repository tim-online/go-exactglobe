package globe

type Costcenters []Costcenter

type Costcenter struct {
	Code string `xml:"code,attr"`

	Description     string               `xml:"Description,omitempty"`
	Description1    string               `xml:"Description1,omitempty"`
	Description2    string               `xml:"Description2,omitempty"`
	Description3    string               `xml:"Description3,omitempty"`
	Description4    string               `xml:"Description4,omitempty"`
	AllocationLevel int                  `xml:"AllocationLevel,omitempty"` // kstpl.ext_dlnivo
	Categories      CostcenterCategories `xml:"Categories,omitempty"`
	GLAccount       *GLAccount           `xml:"GLAccount,omitempty"`
	GLOffset        *GLAccount           `xml:"GLOffset"`
	ExtraFields     ExtraFields          `xml:"ExtraFields,omitempty"`
}
