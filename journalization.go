package globe

type Journalization struct {
	Resource Resource `xml:"Resource"`
	Date     Date     `xml:"Date,omitempty"`
}
