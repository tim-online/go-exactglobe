package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type BankStatement struct {
	// Attributes
	Number string `xml:"number,attr,omitempty"`

	Date              Date               `xml":Date,omitempty`
	GLOffset          GLAccount          `xml:"GLOffset,omitempty"`
	BankStatementLine BankStatementLines `xml:"BankStatementLine,omitempty`
}

func (b BankStatement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(b, e, start)
}

func (b BankStatement) IsEmpty() bool {
	return zero.IsZero(b)
}

type BankStatementLines []BankStatementLine

type BankStatementLine struct {
}
