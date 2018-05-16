package globe

type Addressee struct {
	Name           string `xml:"Name"`
	FirstName      string `xml:"FirstName,omitempty"`
	Initials       string `xml:"Initials,omitempty"`
	Title          Title  `xml:"Title,omitempty"`
	JobDescription string `xml:"JobDescription,omitempty"`
	Email          string `xml:"Email,omitempty"`
}
