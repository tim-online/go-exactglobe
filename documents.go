package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Documents []Document

type Document struct {
}

func (d Document) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(d, e, start)
}

func (d Document) IsEmpty() bool {
	return zero.IsZero(d)
}
