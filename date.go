package globe

import (
	"encoding/xml"
	"time"

	"github.com/cydev/zero"
)

// Date has format YYYY-MM-DD
type Date struct {
	time.Time
}

func (d Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	layout := "2006-01-02"
	return e.EncodeElement(d.Time.Format(layout), start)
	// return omitempty.MarshalXML(d, e, start)
}

func (d Date) IsEmpty() bool {
	return zero.IsZero(d)
}
