package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Projects []Project

type Project struct {
}

func (p Project) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(p, e, start)
}

func (p Project) IsEmpty() bool {
	return zero.IsZero(p)
}
