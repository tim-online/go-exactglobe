package globe

type Payment struct {
	// [ code = string ]
	// [ method = token: { K | B | G | I | A | C | S | W | D | M | V | Y | X | Z } ]
	// [ OriginPackage = token: { F | A | B | C | I | V | T | H | O | K | S | R | W | E | D | P | L | M | Y | U | X | Z } ]>
	// sequence	 	<PaymentMethod> (0,1)
	// <PaymentCondition> (0,1)
	// <Reference> (0,1)
	// <OriginalARAPCurrency> (0,1)
	// <XRateARAPPayment> (0,1)
	// <XRateARAP> (0,1)
	// <TransactionNumberSubAdministration> (0,1)
	// <OriginalARAPAmount> (0,1)
	// <CSSDDate1> (0,1)
	// <CSSDDate2> (0,1)
	// <CSSDYesNo> (0,1)
	// <CSSDAmount1> (0,1)
	// <CSSDAmount2> (0,1)
	// <InvoiceNumber> (0,1)
	// <BankTransactionID> (0,1)
	// <DateLastReminder> (0,1)
	// <BankAccount> (0,1)
	// <InvoiceDueDate> (0,1)
}
