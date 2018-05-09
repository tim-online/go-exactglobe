package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type ExtraFields struct {
	UserFields  UserFields  `xml:"UserFields,omitempty"`
	UserNumbers UserNumbers `xml:"UserNumbers,omitempty"`
}

func (ef *ExtraFields) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if ef == nil {
		return nil
	}
	return omitempty.MarshalXML(ef, e, start)
}

func (ef ExtraFields) IsEmpty() bool {
	return zero.IsZero(ef)
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
