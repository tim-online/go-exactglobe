package globe

type Availability struct {
	DateStart Date `xml:"DateStart"`
	DateEnd   Date `xml:"DateEnd,omitempty"`
}
