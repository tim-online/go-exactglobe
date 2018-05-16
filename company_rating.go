package globe

import "encoding/xml"

type CompanyRating int

func (c CompanyRating) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	tmp := struct {
		XMLName xml.Name `xml:"CompanyRating"`
		Number  int      `xml:"number,attr"`
	}{Number: int(c)}
	return e.EncodeElement(tmp, start)
}
