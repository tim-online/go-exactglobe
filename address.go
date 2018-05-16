package globe

type Addresses []Address

type Address struct {
	// Attributes
	Type string `xml:"type,attr,omitempty"`
	Desc string `xml:"desc,attr,omitempty"`

	Addressee       Addressee       `xml:"Addressee,omitempty"`
	AddressLine1    string          `xml:"AddressLine1,omitempty"`
	AddressLine2    string          `xml:"AddressLine2,omitempty"`
	AddressLine3    string          `xml:"AddressLine3,omitempty"`
	AddressNo       string          `xml:"AddressNo,omitempty"`
	AddressSuffix   string          `xml:"AddressSuffix,omitempty"`
	PostalCode      string          `xml:"PostcalCode,omitempty"`
	City            string          `xml:"City,omitempty"`
	County          string          `xml:"County,omitempty"`
	State           string          `xml:"State,omitempty"`
	Country         string          `xml:"Country"`
	IsoCountry      string          `xml:"IsoCountry,omitempty"`
	IsoRegion       string          `xml:"IsoRegion,omitempty"`
	Phone           string          `xml:"Phone,omitempty"`
	Fax             string          `xml:"Fax,omitempty"`
	VATNumber       string          `xml:"VATNumber,omitempty"`
	Email           string          `xml:"Email,omitempty"`
	Territory       Territory       `xml:"Territory,omitempty"`
	SalesCommission SalesCommission `xml:"SalesCommission,omitempty"`
}
