package globe

type DebtorInvoice struct {
	// Attributes
	Code   string `xml:"code,attr,omitempty"`
	Number int    `xml:"number,attr"`

	Name string `xml:"Name,omitempty"`
}
