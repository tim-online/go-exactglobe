package globe

type PrefixAffix struct {
	Prefix      string `xml:"Prefix,omitempty"`
	Affix       string `xml:"Affix,omitempty"`
	BirthPrefix string `xml:"BirthPrefix,omitempty"`
	BirthAffix  string `xml:"BirthAffix,omitempty"`
}
