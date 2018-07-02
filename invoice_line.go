package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

// docs/XML-Schema.html#ELEMENT_InvoiceLine

type InvoiceLines []InvoiceLine

type InvoiceLine struct {
	Description              string           `xml:"Description,omitempty"`
	Reference1               string           `xml:"Reference1,omitempty"`
	Reference2               string           `xml:"Reference2,omitempty"`
	Reference3               string           `xml:"Reference3,omitempty"`
	YourRef                  string           `xml:"YourRef,omitempty"`
	OurRef                   string           `xml:"OurRef,omitempty"`
	Date                     Date             `xml:"Date,omitempty"`
	Order                    Order            `xml:"Order,omitempty"`
	Currency                 Currency         `xml:"Currency,omitempty"`
	Rate                     float64          `xml:"Rate,omitempty"`
	CalcIncludeVAT           string           `xml:"CalcIncludeVAT,omitempty"`
	AvalaraAddressValidated  string           `xml:"AvalaraAddressValidated,omitempty"`
	UseAvalaraTaxation       string           `xml:"UseAvalaraTaxation,omitempty"`
	Resource                 Resource         `xml:"Resource,omitempty"`
	OrderedBy                OrderedBy        `xml:"OrderedBy,omitempty"`
	OrderedAt                OrderedAt        `xml:"OrderedAt,omitempty"`
	DeliverTo                DeliverTo        `xml:"DeliverTo,omitempty"`
	InvoiceTo                InvoiceTo        `xml:"InvoiceTo,omitempty"`
	Warehouse                Warehouse        `xml:"Warehouse,omitempty"`
	CashRegisterAccount      string           `xml:"CashRegisterAccount,omitempty"`
	PaymentMethod            PaymentMethod    `xml:"PaymentMethod,omitempty"`
	PaymentCondition         PaymentCondition `xml:"PaymentCondition,omitempty"`
	IncoTermCode             IncoTermCode     `xml:IncoTermCode,omitempty`
	IncoTermConfirmPrices    bool             `xml:"IncoTermConfirmPrices,omitempty"`
	IncoTermAcknowledgeOrder bool             `xml:"IncoTermAcknowledgeOrder,omitempty"`
	DeliveryMethod           DeliveryMethod   `xml:"DeliveryMethod,omitempty"`
	Costcenter               Costcenter       `xml:"Costcenter,omitempty"`
	Selection                string           `xml:"Selection,omitempty"`
	ExtraText                ExtraText        `xml:"ExtraText,omitempty"`
	// PurchaseOrderMethod      PurchaseOrderMethod  `xml:"PurchaseOrderMethod,omitempty"`
	// InvoiceGroup          InvoiceGroup         `xml:"InvoiceGroup,omitempty"`
	// Freight               Freight              `xml:"Freight,omitempty"`
	Document           Document           `xml:"Document,omitempty"`
	DocumentAttachment DocumentAttachment `xml:"DocumentAttachment,omitempty"`
	SalesCommission    SalesCommission    `xml:"SalesCommission,omitempty"`
	FreeFields         FreeFields         `xml:"FreeFields,omitempty"`
	CancelDate         Date               `xml:"CancelDate,omitempty"`
	EarlyShipDate      Date               `xml:"EarlyShipDate,omitempty"`
	RequestedShipDate  Date               `xml:"RequestedShipDate,omitempty"`
	BOLNo              string             `xml:"BOLNo,omitempty"`
	PRONo              string             `xml:"PRONo,omitempty"`
	TotalTare          float64            `xml:"TotalTare,omitempty"`
	TotalPacks         float64            `xml:"TotalPacks,omitempty"`
	TotalCartons       float64            `xml:"TotalCartons,omitempty"`
	TotalWeight        float64            `xml:"TotalWeight,omitempty"`
	// ShiptoID              ShiptoID             `xml:"ShiptoID,omitempty"`
	// DistributionCenterID  DistributionCenterID `xml:"DistributionCenterID,omitempty"`
	DestinationName      string `xml:"DestinationName,omitempty"`
	Department           string `xml:"Department,omitempty"`
	ApplyShippingCharges bool   `xml:"ApplyShippingCharges,omitempty"`
	// IsUseAdditionalFields bool                 `xml:"IsUseAdditionalFields,omitempty"`
	// AdditionalFields      AdditionalFields     `xml:"AdditionalFields,omitempty"`
}

func (f InvoiceLine) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(f, e, start)
}

func (f InvoiceLine) IsEmpty() bool {
	return zero.IsZero(f)
}

type OrderedBy struct {
	Debtor  Debtor  `xml:"Debtor"`
	Address Address `xml:"Address,omitempty"`
	Date    Date    `xml:"Date,omitempty"`
}

type OrderedAt struct {
	Creditor Creditor `xml:"Creditor"`
	Address  Address  `xml:"Address,omitempty"`
	Date     Date     `xml:"Date,omitempty"`
}

type DeliverTo struct {
	// Choice: Debtor/Warehouse
	Debtor    Debtor    `xml:"Debtor,omitempty"`
	Warehouse Warehouse `xml:"Warehouse,omitempty"`

	Address Address `xml:"Address,omitempty"`
	Date    Date    `xml:"Date,omitempty"`
}

type InvoiceTo struct {
	// Choice: Debtor/Warehouse
	Debtor    Debtor    `xml:"Debtor,omitempty"`
	Warehouse Warehouse `xml:"Warehouse,omitempty"`

	Address Address `xml:"Address,omitempty"`
	Date    Date    `xml:"Date,omitempty"`
}
