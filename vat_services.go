package globe

type VATServices struct {
	VAT              VAT         `xml:"VAT,omitempty"`
	VATServicesFixed interface{} `xml:"VATServicesFixed"`
}
