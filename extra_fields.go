package globe

type ExtraFields struct {
	UserFields  UserFields  `xml:"UserFields,omitempty"`
	UserNumbers UserNumbers `xml:"UserNumbers,omitempty"`
}

type UserFields []UserField

type UserField struct {
	// Attributes
	Number int `xml:"number,attr"`

	Text string `xml:",chardata"` // @TODO: check if this is correct
}

type UserNumbers []UserNumber

type UserNumber struct {
	// Attributes
	Number int `xml:"number,attr"`

	Double Double `xml:",chardata"` // @TODO: check if this is correct
}
