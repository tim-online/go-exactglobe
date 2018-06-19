package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	errors "github.com/tim-online/go-errors"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type FinReferences struct {
	// Attributes
	TransactionOrigin TransactionOrigin `xml:"TransactionOrigin,attr,omitempty"`

	ProcessNumberJournal  int    `xml:"ProcessNumberJournal,omitempty"`
	UniquePostingNumber   int    `xml:"UniquePostingNumber,omitempty"`
	SequenceNumberEntry   string `xml:"SequenceNumberEntry,omitempty"`
	TransactionNumber2    string `xml:"TransactionNumber2,omitempty"`
	LineCode              string `xml:"LineCode,omitempty"`
	YourRef               string `xml:"YourRef,omitempty"`
	DocumentID            string `xml:"DocumentID,omitempty"`
	DocumentDate          Date   `xml:"DocumentDate,omitempty"`
	DebtorStatementNumber int    `xml:"DebtorStatementNumber,omitempty"`
	StockTrackingNumber   string `xml:"StockTrackingNumber,omitempty"`
	CashRegister          string `xml:"CashRegister,omitempty"`
	ReportDate            Date   `xml:"ReportDate,omitempty"`
}

func (f FinReferences) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(f, e, start)
}

func (f FinReferences) IsEmpty() bool {
	return zero.IsZero(f)
}

type TransactionOrigin string

// I=Invoice, P=Payment, T=Pay in installments, N=None, U=Budget, S=Reconcile
func (to TransactionOrigin) Validate() error {
	valid := []TransactionOrigin{"I", "P", "T", "N", "U", "S"}

	for _, v := range valid {
		if to == v {
			return nil
		}
	}

	return errors.Errorf("Invalid TransactionOrigin '%s'", to)
}
