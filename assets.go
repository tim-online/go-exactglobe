package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Assets []Asset

type Asset struct {
}

func (a Asset) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(a, e, start)
}

func (a Asset) IsEmpty() bool {
	return zero.IsZero(a)
}
