package globe

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"time"

	"github.com/cydev/zero"
)

// Date has format YYYY-MM-DD
type Date struct {
	time.Time
}

func (d Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	layout := "2006-01-02"
	log.Printf("%+v", d.Time)
	return e.EncodeElement(d.Time.Format(layout), start)
	// return omitempty.MarshalXML(d, e, start)
}

func (d Date) IsEmpty() bool {
	return zero.IsZero(d)
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	log.Printf("s: %s", s)
	if err != nil {
		return err
	}

	layout := "2006-01-02"
	d.Time, err = time.Parse(layout, s)
	return err
}
