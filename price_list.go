package globe

type PriceList struct {
	// Attributes
	Code string `xml:"code,attr"`
	Type string `xml:"type,attr,omitempty"` // { S | P }

	// Description          string          `xml:"Description,omitempty"`
	MultiDescriptions       []string     `xml:"Description,omitempty"`
	Currency                Currency     `xml:"Currency,omitempty"`
	Country                 Country      `xml:"Country,omitempty"`
	Availability            Availability `xml:"Availability,omitempty"`
	AccountType             string       `xml:"AccountType,omitempty"`
	AccountStatus           string       `xml:"AccountStatus,omitempty"`
	AccountClassificationId string       `xml:"AccountClassificationId,omitempty"`
}
