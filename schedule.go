package globe

type Schedule struct {
	ID          string `xml:"ID,attr"`
	CompanyCode string `xml:"CompanyCode,attr,omitempty"`
	Weekday     string `xml:"Weekday,attr,omitempty"`
	StartTime   string `xml:"StartTime,attr,omitempty"`
	EndTime     string `xml:"EndTime,attr,omitempty"`
	Syscreator  string `xml:"Syscreator,attr,omitempty"`
}
