package globe

import (
	"encoding/xml"
)

type Double float64

type Bool bool

func (b Bool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if b == true {
		return e.EncodeElement("1", start)
	} else {
		return e.EncodeElement("0", start)
	}
}

func (b *Bool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if *b == true {
		return d.DecodeElement("1", &start)
	} else {
		return d.DecodeElement("0", &start)
	}
}
