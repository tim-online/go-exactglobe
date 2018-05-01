package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

// Date has format YYYY-MM-DD
type Date struct {
}

func (d Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(d, e, start)
}

func (d Date) IsEmpty() bool {
	return zero.IsZero(d)
}
