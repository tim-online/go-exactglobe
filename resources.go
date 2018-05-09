package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Resources []Resource

type Resource struct {
	// Attributes
	Number       int    `xml:"number,attr"`
	Code         string `xml:"code,attr,omitempty"`
	Type         string `xml:"type,attr,omitempty"`   // { A | C | E | R | S | T | V }
	status       string `xml:"status,attr,omitempty"` // { A | H | I | F | R | T }
	Gender       string `xml:"gender,attr,omitempty"`
	License      string `xml:"license,attr,omitempty"`
	AllDivisions int    `xml:"alldivisions,attr,omitempty"`
	Transfer     int    `xml:"transfer,attr,omitempty"`

	LastName   string `xml:"LastName,omitempty"`
	FirstName  string `xml:"FirstName,omitempty"`
	MiddleName string `xml:"MiddleName,omitempty"`
	Initials   string `xml:"Initials,omitempty"`
	Title      string `xml:"Title,omitempty"`
	Image      Image  `xml:"Image,omitempty"`

	MaidenName    string       `xml:"MaidenName,omitempty"`
	Language      Language     `xml:"Language,omitempty"`
	Nationality   Nationality  `xml:"Nationality,omitempty"`
	PrefixAffix   PrefixAffix  `xml:"PrefixAffix,omitempty"`
	OriginCountry string       `xml:"OriginCountry,omitempty"`
	Qualification string       `xml:"Qualification,omitempty"`
	Race          string       `xml:"Race,omitempty"`
	Marital       Marital      `xml:"Marital,omitempty"`
	Note          string       `xml:"Note,omitempty"`
	Office        Office       `xml:"Office,omitempty"`
	Employment    Employment   `xml:"Employment,omitempty"`
	Workschedule  Workschedule `xml:"Workschedule,omitempty"`
	Organization  Organization `xml:"Organization,omitempty"`
	Banking       Banking      `xml:"Banking,omitempty"`
	Contracts     Contracts    `xml:"Contracts,omitempty"`
	HumresLinks   HumresLinks  `xml:"HumresLinks,omitempty"`
	Skills        Skills       `xml:"Skills,omitempty"`
	Territories   Territories  `xml:"Territories,omitempty"`
	FreeFields    FreeFields   `xml:"FreeFields,omitempty"`
}

func (r Resource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(r, e, start)
}

func (r Resource) IsEmpty() bool {
	return zero.IsZero(r)
}
