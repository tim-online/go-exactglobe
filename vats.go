package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type VATs []VAT

type VAT struct {
}

func (v VAT) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(v, e, start)
}

func (v VAT) IsEmpty() bool {
	return zero.IsZero(v)
}
