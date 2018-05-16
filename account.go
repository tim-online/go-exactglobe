package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	errors "github.com/tim-online/go-errors"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Accounts []Account

type Account struct {
	// Attributes
	Code     string          `xml:"code,attr"`
	ID       string          `xml:"ID,attr,omitempty"`
	Type     AccountType     `xml:"type,attr,omitempty"`     // { A | B | C | D | E | L | N | P | R | S | T }
	Status   AccountStatus   `xml:"status,attr,omitempty"`   // { A | B | E | N | P | R | S }
	CodeType AccountCodeType `xml:"codetype,attr,omitempty"` // { C | D | A }

	Name                     string          `xml:"Name,omitempty"`
	Phone                    string          `xml:"Phone,omitempty"`
	PhoneExt                 string          `xml:"PhoneExt,omitempty"`
	Fax                      string          `xml:"Fax,omitempty"`
	Email                    string          `xml:"Email,omitempty"`
	HomePage                 string          `xml:"HomePage,omitempty"`
	Addresses                Addresses       `xml:"Addresses>Address,omitempty"`
	Contacts                 Contacts        `xml:"Contacts>Contact,omitempty"`
	Manager                  Resource        `xml:"Manager,omitempty"`
	Parent                   *Account        `xml:"Parent,omitempty"`
	Image                    Image           `xml:"Image,omitempty"`
	Note                     string          `xml:"Note,omitempty"`
	AccountCode              string          `xml:"AccountCode,omitempty"`
	Debtor                   Debtor          `xml:"Debtor,omitempty"`
	Creditor                 Creditor        `xml:"Creditor,omitempty"`
	VAT                      VAT             `xml:"VAT,omitempty"`
	VATFixed                 bool            `xml:"VATFixed,omitempty"`
	VATServices              VATServices     `xml:"VATServices,omitempty"`
	VATNumber                string          `xml:"VATNumber,omitempty"`
	VATCheckDate             Date            `xml:"VATCheckDate,omitempty"`
	VATLiability             string          `xml:"VATLiability,omitempty"`
	FiscalCode               string          `xml:"FiscalCode,omitempty"`
	IsTaxExempted            bool            `xml:"IsTaxExempted,omitempty"`
	ExemptNumber             string          `xml:"ExemptNumber,omitempty"`
	ShowAttention            bool            `xml:"ShowAttention,omitempty"`
	ChamberOfCommerce        string          `xml:"ChamberOfCommerce,omitempty"`
	FederalCategory          string          `xml:"FederalCategory,omitempty"`
	FederalIDNumber          string          `xml:"FederalIDNumber,omitempty"`
	FederalIDType            string          `xml:"FederalIDType,omitempty"`
	PaymentMethod            PaymentMethod   `xml:"PaymentMethod,omitempty"`
	PaymentCondition         string          `xml:"PaymentCondition,omitempty"`
	DeliveryMethod           DeliveryMethod  `xml:"DeliveryMethod,omitempty"`
	IncoTermCode             IncoTermCode    `xml:"IncoTermCode,omitempty"`
	IncoTermConfirmPrices    bool            `xml:"IncoTermConfirmPrices,omitempty"`
	IncoTermAcknowledgeOrder bool            `xml:"IncoTermAcknowledgeOrder,omitempty"`
	PriceList                PriceList       `xml:"PriceList,omitempty"`
	Selection                Selection       `xml:"Selection,omitempty"`
	ExtraText                ExtraText       `xml:"ExtraText,omitempty"`
	CompanySize              CompanySize     `xml:"CompanySize,omitempty"`
	CompanyOrigin            CompanyOrigin   `xml:"CompanyOrigin,omitempty"`
	Sector                   Sector          `xml:"Sector,omitempty"`
	AccountCategory          AccountCategory `xml:"AccountCategory,omitempty"`
	CompanyRating            CompanyRating   `xml:"CompanyRating,omitempty"`
	Categories               Categories      `xml:"Categories,omitempty"`
	PayeeName                string          `xml:"PayeeName,omitempty"`
	DunsNumber               string          `xml:"DunsNumber,omitempty"`
	FreeFields               FreeFields      `xml:"FreeFields,omitempty"`
	GLN                      string          `xml:"GLN,omitempty"`
}

func (a Account) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(a, e, start)
}

func (a Account) IsEmpty() bool {
	return zero.IsZero(a)
}

type AccountType string

// A=Associate, B=Bank, C=Customer, D=Division, E=Employee, L=Lead, N=Not validated, P=Prospect, R=Reseller, S=Supplier, T=Suspect
func (a AccountType) Validate() error {
	valid := []AccountType{"A", "B", "C", "D", "E", "L", "N", "P", "R", "S", "T"}

	for _, v := range valid {
		if a == v {
			return nil
		}
	}

	return errors.Errorf("Invalid AccountType '%s'", a)
}

type AccountStatus string

// A=Active, B=Blocked, E=Exit, N=Not validated, P=Pilot, R=Reference, S=Passive
func (a AccountStatus) Validate() error {
	valid := []AccountStatus{"A", "B", "E", "N", "P", "R", "S"}

	for _, v := range valid {
		if a == v {
			return nil
		}
	}

	return errors.Errorf("Invalid AccountStatus '%s'", a)
}

type AccountCodeType string

// C=Creditor code, D=Debtor code, A=Account code only
func (a AccountCodeType) Validate() error {
	valid := []AccountCodeType{"C", "D", "A"}

	for _, v := range valid {
		if a == v {
			return nil
		}
	}

	return errors.Errorf("Invalid AccountCodeType '%s'", a)
}
