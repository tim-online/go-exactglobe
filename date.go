package globe

import (
	"encoding/json"
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
}

func (d Date) IsEmpty() bool {
	return zero.IsZero(d)
}

func (d Date) MarshalJSON() ([]byte, error) {
	layout := "2006-01-02"
	return json.Marshal(d.Time.Format(layout))
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	layout := "2006-01-02"
	d.Time, err = time.Parse(layout, s)
	return err
}
