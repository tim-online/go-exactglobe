package globe

type FreeFields struct {
	FreeTexts   FreeTexts   `xml:"FreeTexts,omitempty"`
	FreeDates   FreeDates   `xml:"FreeDates,omitempty"`
	FreeNumbers FreeNumbers `xml:"FreeNumbers,omitempty"`
	FreeYesNos  FreeYesNos  `xml:"FreeYesNos,omitempty"`
	FreeOptions FreeOptions `xml:"FreeOptions,omitempty"`
}

type FreeTexts []FreeText

type FreeText struct {
	// Attributes
	Number int    `xml:"number,attr"`
	Label  string `xml:"label,attr,omitempty"`

	Text string `xml:",chardata"` // @TODO: check if this is correct
}

type FreeDates []FreeDate

type FreeDate struct {
	Number int    `xml:"number,attr"`
	Label  string `xml:"label,attr,omitempty"`

	Date Date `xml:",chardata"` // @TODO: check if this is correct
}

type FreeNumbers []FreeNumber

type FreeNumber struct {
	// Attributes
	Number int    `xml:"number,attr"`
	Label  string `xml:"label,attr,omitempty"`

	Float float64 `xml:",chardata"` // @TODO: check if this is correct
}

type FreeYesNos []FreeYesNo

type FreeYesNo struct {
	// Attributes
	YesNo int    `xml:"number,attr"`
	Label string `xml:"label,attr,omitempty"`

	Bool bool `xml:",chardata"` // @TODO: check if this is correct
}

type FreeOptions []FreeOption

type FreeOption = FreeYesNo
