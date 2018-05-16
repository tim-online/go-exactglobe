package globe

type JobActivity struct {
	// Attributes
	Code string `xml:"code,attr,omitempty"`

	Description string   `xml:"Description,omitempty"`
	JobGroup    JobGroup `xml:"JobGroup,omitempty"`
}
