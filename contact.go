package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	errors "github.com/tim-online/go-errors"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Contacts []Contact

type Contact struct {
	// Attributes
	ID      string        `xml:"ID,attr,omitempty"`
	Gender  string        `xml:"gender,attr,omitempty"`
	Default bool          `xml:"default,attr,omitempty"`
	Status  ContactStatus `xml:"status,attr,omitempty"`

	LastName       string     `xml:"LastName,omitempty"`
	FirstName      string     `xml:"FirstName,omitempty"`
	MiddleName     string     `xml:"MiddleName,omitempty"`
	Initials       string     `xml:"Initials,omitempty"`
	Title          Title      `xml:"Title,omitempty"`
	Image          Image      `xml:"Image,omitempty"`
	Addresses      Addresses  `xml:"Addresses,omitempty"`
	Manager        Resource   `xml:"Manager,omitempty"`
	Language       Language   `xml:"Language,omitempty"`
	JobTitle       JobTitle   `xml:"JobTitle,omitempty"`
	JobDescription string     `xml:"JobDescription,omitempty"`
	Phone          string     `xml:"Phone,omitempty"`
	PhoneExt       string     `xml:"PhoneExt,omitempty"`
	Fax            string     `xml:"Fax,omitempty"`
	Mobile         string     `xml:"Mobile,omitempty"`
	MobileShortcut string     `xml:"MobileShortcut,omitempty"`
	Email          string     `xml:"Email,omitempty"`
	WebAccess      bool       `xml:"WebAccess,omitempty"`
	Note           string     `xml:"Note,omitempty"`
	FreeFields     FreeFields `xml:"FreeFields,omitempty"`
}

func (c Contact) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c Contact) IsEmpty() bool {
	return zero.IsZero(c)
}

// A=Active, E=Exit
type ContactStatus string

func (c ContactStatus) Validate() error {
	valid := []ContactStatus{"A", "E"}

	if c == "" {
		return nil
	}

	for _, v := range valid {
		if c == v {
			return nil
		}
	}

	return errors.Errorf("Invalid EntryLineType '%s'", c)
}
