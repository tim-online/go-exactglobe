package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type FinReferences struct {
	// [ TransactionOrigin = token: { B | I | P | T | N | U | S } ]>
	// sequence	 	<ProcessNumberJournal> (0,1)
	// <UniquePostingNumber> (0,1)
	// <SequenceNumberEntry> (0,1)
	// <TransactionNumber2> (0,1)
	// <LineCode> (0,1)
	// <YourRef> (0,1)
	// <DocumentID> (0,1)
	// <DocumentDate> (0,1)
	// <DebtorStatementNumber> (0,1)
	// <StockTrackingNumber> (0,1)
	// <CashRegister> (0,1)
	// <ReportDate> (0,1)
}

func (f FinReferences) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(f, e, start)
}

func (f FinReferences) IsEmpty() bool {
	return zero.IsZero(f)
}
