package globe

type CompanyOrigin struct {
	// Attributes
	Code string `xml:"code,attr"`

	Description string `xml:"Description,omitempty"`
}
