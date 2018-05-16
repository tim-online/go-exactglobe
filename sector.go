package globe

type Sector struct {
	// Attributes
	Code string `xml:"code,attr"`

	Description       string            `xml:"Description,omitempty"`
	MultiDescriptions MultiDescriptions `xml:"MultiDescriptions,omitempty"`
	SubSector         SubSector         `xml:"SubSector,omitempty"`
}
