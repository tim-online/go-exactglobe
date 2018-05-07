package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	errors "github.com/tim-online/go-errors"
	"github.com/tim-online/go-exactglobe/omitempty"
)

// docs/XML-Schema.html#ELEMENT_FinEntryLine

type FinEntryLine struct {
	Number          int              `xml:"number,attr,omitempty"`
	Type            FinEntryLineType `xml:"type,attr,omitempty"`    // { B | C | E | F | I | N | O | P | V | X }
	Subype          string           `xml:"subtype,attr,omitempty"` // { N | K | C | T | Q | Z | Y | R | P | S | D | F | M | L | E | I | V | A | B | G | H | J | W }
	Transactiontype int              `xml:"transactiontype,attr,omitempty"`
	Linecode        string           `xml:"linecode,attr,omitempty"` // { A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | X }
	Code            string           `xml:"code,attr,omitempty"`
	ID              string           `xml:"ID,attr,omitempty"`
	Correction      bool             `xml:"correction,attr,omitempty"`

	Date                  Date               `xml:"Date"`
	Starttime             Date               `xml:"Starttime,omitempty"`
	Endtime               Date               `xml:"Endtime,omitempty"`
	FinYear               FinYear            `xml:"FinYear,omitempty"`
	FinPeriod             FinPeriod          `xml:"FinPeriod,omitempty"`
	GLAccount             GLAccount          `xml:"GLAccount,omitempty"`
	Description           string             `xml:"Description,omitempty"`
	Costcenter            Costcenter         `xml:"Costcenter,omitempty"`
	Costunit              Costunit           `xml:"Costunit,omitempty"`
	Debtor                Debtor             `xml:"Debtor,omitempty"`
	Creditor              Creditor           `xml:"Creditor,omitempty"`
	Country               Country            `xml:"Country,omitempty"`
	Resource              Resource           `xml:"Resource,omitempty"`
	Item                  Item               `xml:"Item,omitempty"`
	Asset                 Asset              `xml:"Asset,omitempty"`
	Warehouse             Warehouse          `xml:"Warehouse,omitempty"`
	WarehouseLocation     WarehouseLocation  `xml:"WarehouseLocation,omitempty"`
	Project               Project            `xml:"Project,omitempty"`
	Quantity              float64            `xml:"Quantity,omitempty"`
	Amount                float64            `xml:"Amount,omitempty"`
	ForeignAmount         ForeignAmount      `xml:"ForeignAmount,omitempty"`
	OffsetGL              string             `xml:"OffsetGL,omitempty"`
	TransactionType       int                `xml:"TransactionType,omitempty"` // gbkmut.transactiontype (tag goes obsolete, is replaced by attribute in FinEntryLine)
	TransactionNumber     string             `xml:"TransactionNumber,omitempty"`
	VATTransaction        VATTransaction     `xml:"VATTransaction,omitempty"`
	VATTransaction2       VATTransaction     `xml:"VATTransaction2,omitempty"`
	VATTransaction3       VATTransaction     `xml:"VATTransaction3,omitempty"`
	VATTransaction4       VATTransaction     `xml:"VATTransaction4,omitempty"`
	VATTransaction5       VATTransaction     `xml:"VATTransaction5,omitempty"`
	OrderDebtor           string             `xml:"OrderDebtor,omitempty"`
	Payment               Payment            `xml:"Payment,omitempty"`
	Delivery              Delivery           `xml:"Delivery,omitempty"`
	Reminder              bool               `xml:"Reminder,omitempty"`
	FinReferences         FinReferences      `xml:"FinReferences,omitempty"`
	Document              Document           `xml:"Document,omitempty"`
	DocumentAttachment    DocumentAttachment `xml:"DocumentAttachment,omitempty"`
	Selection             Selection          `xml:"Selection,omitempty"`
	Unit                  Unit               `xml:"Unit,omitempty"`
	PriceList             PriceList          `xml:"PriceList,omitempty"`
	Discount              float64            `xml:"discount,omitempty"`
	IntrastatStandardCode string             `xml:"IntrastatStandardCode,omitempty"`
	FreeFields            FreeFields         `xml:"FreeFields,omitempty"`
}

func (f FinEntryLine) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(f, e, start)
}

func (f FinEntryLine) IsEmpty() bool {
	return zero.IsZero(f)
}

// B=Budget, C=Balance Correction, E=Elimination, F=Fiscal, I=Intercompany, N=Normal (default), O=Obligation, P=Opening balance correction, V=Void, X=Non ledger
type FinEntryLineType string

func (f FinEntryLineType) Validate() error {
	valid := []FinEntryLineType{"B", "C", "E", "F", "I", "N", "O", "P", "V", "X"}

	if f == "" {
		return nil
	}

	for _, v := range valid {
		if f == v {
			return nil
		}
	}

	return errors.Errorf("Invalid EntryLineType '%s'", f)
}
