package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	errors "github.com/tim-online/go-errors"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Payment struct {
	// Attributes
	Code          string               `xml:"code,attr,omitempty"`
	Method        Paymentmethod        `xml:"method,attr,omitempty"`
	OriginPackage PaymentOriginPackage `xml:"OriginPackage,attr,omitempty"`

	Paymentmethod                      Paymentmethod    `xml:"Paymentmethod,omitempty"`
	PaymentCondition                   PaymentCondition `xml:"PaymentCondition,omitempty"`
	Reference                          string           `xml:"Reference,omitempty"`
	OriginalARAPCurrency               string           `xml:"OriginalARAPCurrency,omitempty"`
	XRateARAPPayment                   float64          `xml:"XRateARAPPayment,omitempty"`
	XRateARAP                          float64          `xml:"XRateARAP,omitempty"`
	TransactionNumberSubAdministration string           `xml:"TransactionNumberSubAdministration,omitempty"`
	OriginalARAPAmount                 float64          `xml:"OriginalARAPAmount,omitempty"`
	CSSDDate1                          Date             `xml:"CSSDDate1,omitempty"`
	CSSDDate2                          Date             `xml:"CSSDDate2,omitempty"`
	CSSDYesNo                          string           `xml:"CSSDYesNo,omitempty"`
	CSSDAmount1                        float64          `xml:"CSSDAmount1,omitempty"`
	CSSDAmount2                        float64          `xml:"CSSDAmount2,omitempty"`
	InvoiceNumber                      int              `xml:"InvoiceNumber,omitempty"`
	BankTransactionID                  string           `xml:"BankTransactionID,omitempty"`
	DateLastReminder                   Date             `xml:"DateLastReminder,omitempty"`
	BankAccount                        BankAccount      `xml:"BankAccount,omitempty"`
	InvoiceDueDate                     Date             `xml:"InvoiceDueDate,omitempty"`
}

func (p Payment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(p, e, start)
}

func (p Payment) IsEmpty() bool {
	return zero.IsZero(p)
}

// K=Cash, B=Bank, G=Giro, I=Collection, A=Giro collection slip, S=Bacs, W=Bill of exchange, D=Bill of exchange, M=Domiciliation, V=ESR payments, Y=ES payments, X=Payments in CHF and FC, Z=Bank cheques
type Paymentmethod string

func (p Paymentmethod) Validate() error {
	valid := []Paymentmethod{"K", "B", "G", "I", "A", "C", "S", "W", "D", "M", "V", "Y", "X", "Z"}

	if p == "" {
		return nil
	}

	for _, v := range valid {
		if p == v {
			return nil
		}
	}

	return errors.Errorf("Invalid Paymentmethod '%s'", p)
}

// F=E-Invoice, A=E-Account, B=E-Payments, C=E-Bank, I=E-Collection, V=E-Assets, T=recurring entries, H=revaluation, O=opening new FY, K=E-Column, S=E-Cost Allocation, R=E-Stock Purchase, W=B/E accounts, E=IncInvReg., D=closing entry, P=E-Job Cost., L=E-Service Management, M=E-PAS, Y=E-Payroll , U=Budget, X=XML
type PaymentOriginPackage string

func (p PaymentOriginPackage) Validate() error {
	valid := []PaymentOriginPackage{"F", "A", "B", "C", "I", "V", "T", "H", "O", "K", "S", "R", "W", "E", "D", "P", "L", "M", "Y", "U", "X", "Z"}

	if p == "" {
		return nil
	}

	for _, v := range valid {
		if p == v {
			return nil
		}
	}

	return errors.Errorf("Invalid PaymentOriginPackage '%s'", p)
}
