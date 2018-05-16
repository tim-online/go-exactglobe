package globe

type JobTitle struct {
	// Attributes
	Code string `xml:"code,attr"`

	Description  string      `xml:"Description,omitempty"`
	Level        int         `xml:"Level,omitempty"`
	JobGroup     JobGroup    `xml:"JobGroup,omitempty"`
	JobActivity  JobActivity `xml:"JobActivity,omitempty"`
	IsProductive bool        `xml:"IsProductive,omitempty"`
	ScaleMin     int         `xml:"ScaleMin,omitempty"`
	ScaleMax     int         `xml:"ScaleMax,omitempty"`
}
