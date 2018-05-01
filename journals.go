package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Journals []Journal

type Journal struct {
	// Attributes
	Code string `xml:"code,attr"`
	Type string `xml:"type,attr"` // { B | G | K | M | I | T | V }

	// Description          string          `xml:"Description,omitempty"`
	MultiDescriptions    []string        `xml:"Description,omitempty"`
	Abbreviation         string          `xml:"Abbreviation,omitempty"`
	Currency             Currency        `xml:"Currency,omitempty"`
	GLAccount            GLAccount       `xml:"GLAccount,omitempty"`
	GLPaymentInTransit   GLAccount       `xml:"GLPaymentInTransit,omitempty"`
	FreeFields           FreeFields      `xml:"FreeFields,omitempty"`
	JournalSettings      JournalSettings `xml:"JournalSettings,omitempty"`
	VariableCurrency     interface{}     `xml:"VariableCurrency,omitempty"`
	VariableExRate       interface{}     `xml:"VariableExRate,omitempty"`
	ExitAutomatically    interface{}     `xml:"ExitAutomatically,omitempty"`
	BlockOutstandingItem interface{}     `xml:"BlockOutstandingItem,omitempty"`
	UseIntercompany      interface{}     `xml:"UseIntercompany,omitempty"`
}

func (j Journal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(j, e, start)
}

func (j Journal) IsEmpty() bool {
	return zero.IsZero(j)
}

type JournalSettings []JournalSetting

type JournalSetting struct {
	FinYear             FinYear `xml:"FinYear"`
	EntryNumber         string  `xml:"EntryNumber"`
	UniquePostingNumber int     `xml:"UniquePostingNumber"`
}
