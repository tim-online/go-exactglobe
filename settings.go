package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Settings struct {
	SettingGroup SettingGroup `xml:">SettingGroup"`
}

func (s Settings) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(s, e, start)
}

func (s Settings) IsEmpty() bool {
	return zero.IsZero(s)
}
