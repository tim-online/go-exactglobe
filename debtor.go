package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

// The maximum length of the number attribute is 6. If code is filled, then
// number can be left empty.
type Debtor struct {
	// Attributes
	Code   string `xml:"code,attr"`
	Number string `xml:"number,attr,omitempty"`
	Type   string `xml:"type,attr,omitempty"` // { A | B | C | D | E | L | N | P | R | S | T }

	Name                   string          `xml:"Name,omitempty"`
	Currency               Currency        `xml:"Currency,omitempty"`
	SecurityLevel          int             `xml:"SecurityLevel,omitempty"`
	BankAccounts           BankAccounts    `xml:"BankAccounts,omitempty"`
	ItemAccounts           ItemAccounts    `xml:"ItemAccounts,omitempty"`
	AutoMatching           bool            `xml:"AutoMatching,omitempty"`
	GLOffset               GLAccount       `xml:"GLOffset,omitempty"`
	GLCentralization       GLAccount       `xml:"GLCentralization,omitempty"`
	ExternalCode           string          `xml:"ExternalCode,omitempty"`
	DebtorInvoice          DebtorInvoice   `xml:"DebtorInvoice,omitempty"`
	CreditLine             float64         `xml:"CreditLine,omitempty"`
	Discount               float64         `xml:"Discount,omitempty"`
	SendReminder           Bool            `xml:"SendReminder,omitempty"`
	AccountEmployee        string          `xml:"AccountEmployee,omitempty"`
	DateLastReminder       Date            `xml:"DateLastReminder,omitempty"`
	PrintStatement         bool            `xml:"PrintStatement,omitempty"`
	OrderConfirmation      string          `xml:"OrderConfirmation,omitempty"`
	AllowPartialShipment   bool            `xml:"AllowPartialShipment,omitempty"`
	AddExtraReceiptToOrder bool            `xml:"AddExtraReceiptToOrder,omitempty"`
	AllowBackOrders        bool            `xml:"AllowBackOrders,omitempty"`
	InvoiceCopies          int             `xml:"InvoiceCopies,omitempty"`
	InvoiceMethod          string          `xml:"InvoiceMethod,omitempty"`
	GroupInvoice           string          `xml:"GroupInvoice,omitempty"`
	AttachUBL              bool            `xml:"AttachUBL,omitempty"`
	Mailbox                string          `xml:"Mailbox,omitempty"`
	IsCommissionable       bool            `xml:"IsCommissionable,omitempty"`
	Territory              Territory       `xml:"Territory,omitempty"`
	SalesCommission        SalesCommission `xml:"SalesCommission,omitempty"`
	ApplyShippingCharges   bool            `xml:"ApplyShippingCharges,omitempty"`
}

func (d Debtor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(d, e, start)
}

func (d Debtor) IsEmpty() bool {
	return zero.IsZero(d)
}
